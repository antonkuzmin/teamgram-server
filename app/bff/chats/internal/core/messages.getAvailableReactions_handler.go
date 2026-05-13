package core

import (
	"github.com/teamgram/proto/mtproto"
)

var defaultAvailableReactions = []*mtproto.AvailableReaction{
	{Reaction: "👍", Title: "👍"},
	{Reaction: "👎", Title: "👎"},
	{Reaction: "❤️", Title: "❤️"},
	{Reaction: "🔥", Title: "🔥"},
	{Reaction: "🥰", Title: "🥰"},
	{Reaction: "👏", Title: "👏"},
	{Reaction: "😁", Title: "😁"},
	{Reaction: "🤔", Title: "🤔"},
	{Reaction: "🤯", Title: "🤯"},
	{Reaction: "😱", Title: "😱"},
	{Reaction: "🤬", Title: "🤬"},
	{Reaction: "😢", Title: "😢"},
	{Reaction: "🎉", Title: "🎉"},
	{Reaction: "🤩", Title: "🤩"},
	{Reaction: "🤮", Title: "🤮"},
	{Reaction: "💩", Title: "💩"},
	{Reaction: "🙏", Title: "🙏"},
	{Reaction: "👌", Title: "👌"},
	{Reaction: "🕊️", Title: "🕊️"},
	{Reaction: "🤡", Title: "🤡"},
	{Reaction: "🥱", Title: "🥱"},
	{Reaction: "🥴", Title: "🥴"},
	{Reaction: "😍", Title: "😍"},
	{Reaction: "🐳", Title: "🐳"},
	{Reaction: "🌚", Title: "🌚"},
	{Reaction: "🌭", Title: "🌭"},
	{Reaction: "💯", Title: "💯"},
	{Reaction: "🤣", Title: "🤣"},
	{Reaction: "⚡️", Title: "⚡️"},
	{Reaction: "🍌", Title: "🍌"},
	{Reaction: "🏆", Title: "🏆"},
	{Reaction: "💔", Title: "💔"},
	{Reaction: "🤨", Title: "🤨"},
	{Reaction: "😐", Title: "😐"},
	{Reaction: "🍓", Title: "🍓"},
	{Reaction: "🍾", Title: "🍾"},
	{Reaction: "💋", Title: "💋"},
	{Reaction: "🖕", Title: "🖕"},
	{Reaction: "😈", Title: "😈"},
	{Reaction: "😴", Title: "😴"},
	{Reaction: "😭", Title: "😭"},
	{Reaction: "🤓", Title: "🤓"},
	{Reaction: "👻", Title: "👻"},
	{Reaction: "👨‍💻", Title: "👨‍💻"},
	{Reaction: "👀", Title: "👀"},
	{Reaction: "🎃", Title: "🎃"},
	{Reaction: "🙈", Title: "🙈"},
	{Reaction: "😇", Title: "😇"},
	{Reaction: "😨", Title: "😨"},
	{Reaction: "🤝", Title: "🤝"},
	{Reaction: "✍️", Title: "✍️"},
	{Reaction: "🤗", Title: "🤗"},
	{Reaction: "🫡", Title: "🫡"},
	{Reaction: "🎅", Title: "🎅"},
	{Reaction: "🎄", Title: "🎄"},
	{Reaction: "☃️", Title: "☃️"},
	{Reaction: "💅", Title: "💅"},
	{Reaction: "🤪", Title: "🤪"},
	{Reaction: "🗿", Title: "🗿"},
	{Reaction: "🆒", Title: "🆒"},
	{Reaction: "💘", Title: "💘"},
	{Reaction: "🙉", Title: "🙉"},
	{Reaction: "🦄", Title: "🦄"},
	{Reaction: "😘", Title: "😘"},
	{Reaction: "💊", Title: "💊"},
	{Reaction: "🙊", Title: "🙊"},
	{Reaction: "😎", Title: "😎"},
	{Reaction: "👾", Title: "👾"},
	{Reaction: "🤷‍♂️", Title: "🤷‍♂️"},
	{Reaction: "🤷", Title: "🤷"},
	{Reaction: "🤷‍♀️", Title: "🤷‍♀️"},
	{Reaction: "😡", Title: "😡"},
}

func (c *ChatsCore) MessagesGetAvailableReactions(in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	emptyDoc := mtproto.MakeTLDocumentEmpty(nil).To_Document()

	reactions := make([]*mtproto.AvailableReaction, 0, len(defaultAvailableReactions))
	for _, r := range defaultAvailableReactions {
		reactions = append(reactions, mtproto.MakeTLAvailableReaction(&mtproto.AvailableReaction{
			Reaction:          r.Reaction,
			Title:             r.Title,
			StaticIcon:        emptyDoc,
			AppearAnimation:   emptyDoc,
			SelectAnimation:   emptyDoc,
			ActivateAnimation: emptyDoc,
			EffectAnimation:   emptyDoc,
		}).To_AvailableReaction())
	}

	return mtproto.MakeTLMessagesAvailableReactions(&mtproto.Messages_AvailableReactions{
		Hash:      int32(in.GetHash()),
		Reactions: reactions,
	}).To_Messages_AvailableReactions(), nil
}
