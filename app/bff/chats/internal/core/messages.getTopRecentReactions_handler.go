package core

import (
	"github.com/teamgram/proto/mtproto"
)

var topReactions = []*mtproto.Reaction{
	mtproto.MakeTLReactionEmoji(&mtproto.Reaction{Emoticon: "👍"}).To_Reaction(),
	mtproto.MakeTLReactionEmoji(&mtproto.Reaction{Emoticon: "❤️"}).To_Reaction(),
	mtproto.MakeTLReactionEmoji(&mtproto.Reaction{Emoticon: "🔥"}).To_Reaction(),
	mtproto.MakeTLReactionEmoji(&mtproto.Reaction{Emoticon: "😁"}).To_Reaction(),
	mtproto.MakeTLReactionEmoji(&mtproto.Reaction{Emoticon: "🥰"}).To_Reaction(),
}

func (c *ChatsCore) MessagesGetTopReactions(in *mtproto.TLMessagesGetTopReactions) (*mtproto.Messages_Reactions, error) {
	limit := int(in.GetLimit())
	if limit <= 0 || limit > len(topReactions) {
		limit = len(topReactions)
	}
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Hash:      in.GetHash(),
		Reactions: topReactions[:limit],
	}).To_Messages_Reactions(), nil
}

func (c *ChatsCore) MessagesGetRecentReactions(in *mtproto.TLMessagesGetRecentReactions) (*mtproto.Messages_Reactions, error) {
	limit := int(in.GetLimit())
	if limit <= 0 || limit > len(topReactions) {
		limit = len(topReactions)
	}
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Hash:      in.GetHash(),
		Reactions: topReactions[:limit],
	}).To_Messages_Reactions(), nil
}

func (c *ChatsCore) MessagesGetDefaultTagReactions(in *mtproto.TLMessagesGetDefaultTagReactions) (*mtproto.Messages_Reactions, error) {
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Hash:      in.GetHash(),
		Reactions: topReactions,
	}).To_Messages_Reactions(), nil
}
