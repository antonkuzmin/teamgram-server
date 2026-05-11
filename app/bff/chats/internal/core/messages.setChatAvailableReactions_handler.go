package core

import (
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

func (c *ChatsCore) MessagesSetChatAvailableReactions(in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	var (
		peer     = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		rUpdates *mtproto.Updates
		date     = time.Now().Unix()
	)

	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		var (
			availableReactionsType int32
			availableReactions     []string
		)

		if chatReactions := in.GetAvailableReactions_CHATREACTIONS(); chatReactions != nil {
			availableReactionsType, availableReactions = chatReactions.ToChatReactions()
		} else if vectorString := in.GetAvailableReactions_VECTORSTRING(); len(vectorString) > 0 {
			availableReactionsType = mtproto.ChatReactionsTypeSome
			availableReactions = vectorString
		} else {
			availableReactionsType = mtproto.ChatReactionsTypeNone
		}

		chat, err := c.svcCtx.Dao.ChatClient.Client().ChatSetChatAvailableReactions(c.ctx, &chatpb.TLChatSetChatAvailableReactions{
			SelfId:                 c.MD.UserId,
			ChatId:                 peer.PeerId,
			AvailableReactionsType: availableReactionsType,
			AvailableReactions:     availableReactions,
		})
		if err != nil {
			c.Logger.Errorf("messages.setChatAvailableReactions - error: %v", err)
			return nil, err
		}

		updates := mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
			Date:    int32(date),
			Seq:     0,
		}).To_Updates()

		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
			UserId:        c.MD.UserId,
			PermAuthKeyId: c.MD.PermAuthKeyId,
			Updates:       updates,
		})

		chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
			if userId != c.MD.UserId {
				c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
					UserId:  userId,
					Updates: updates,
				})
			}
			return nil
		})

		rUpdates = updates
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.setChatAvailableReactions blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("invalid peer type: {%v}")
		return nil, err
	}

	return rUpdates, nil
}
