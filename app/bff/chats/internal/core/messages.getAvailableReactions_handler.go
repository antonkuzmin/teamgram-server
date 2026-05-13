package core

import (
	"github.com/teamgram/proto/mtproto"
)

func (c *ChatsCore) MessagesGetAvailableReactions(in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	reactions := make([]*mtproto.AvailableReaction, 0, len(defaultAvailableReactions))
	for _, def := range defaultAvailableReactions {
		doc := emojiDocuments[def.Emoji]
		if doc == nil {
			continue
		}
		reactions = append(reactions, mtproto.MakeTLAvailableReaction(&mtproto.AvailableReaction{
			Reaction:          def.Emoji,
			Title:             def.Emoji,
			StaticIcon:        doc,
			AppearAnimation:   doc,
			SelectAnimation:   doc,
			ActivateAnimation: doc,
			EffectAnimation:   doc,
		}).To_AvailableReaction())
	}

	return mtproto.MakeTLMessagesAvailableReactions(&mtproto.Messages_AvailableReactions{
		Hash:      int32(in.GetHash()),
		Reactions: reactions,
	}).To_Messages_AvailableReactions(), nil
}
