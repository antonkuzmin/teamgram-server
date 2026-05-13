package core

import (
	"os"
	"path/filepath"

	"github.com/teamgram/proto/mtproto"
)

var emojiDocuments map[string]*mtproto.Document

func init() {
	emojiDocuments = make(map[string]*mtproto.Document, len(defaultAvailableReactions))
	for i, r := range defaultAvailableReactions {
		id := int64(0x10000000 + i)
		doc, err := buildEmojiDocument(id, r)
		if err != nil {
			continue
		}
		emojiDocuments[r.Emoji] = doc
	}
}

type emojiDef struct {
	Emoji    string
	Codepoint string
}

var defaultAvailableReactions = []emojiDef{
	{"👍", "1f44d"},
	{"👎", "1f44e"},
	{"❤️", "2764"},
	{"🔥", "1f525"},
	{"🥰", "1f970"},
	{"👏", "1f44f"},
	{"😁", "1f601"},
	{"🤔", "1f914"},
	{"🤯", "1f92f"},
	{"😱", "1f631"},
	{"🤬", "1f92c"},
	{"😢", "1f622"},
	{"🎉", "1f389"},
	{"🤩", "1f929"},
	{"🤮", "1f92e"},
	{"💩", "1f4a9"},
	{"🙏", "1f64f"},
	{"👌", "1f44c"},
	{"🕊️", "1f54a"},
	{"🤡", "1f921"},
	{"🥱", "1f971"},
	{"🥴", "1f974"},
	{"😍", "1f60d"},
	{"🐳", "1f433"},
	{"🌚", "1f31a"},
	{"🌭", "1f32d"},
	{"💯", "1f4af"},
	{"🤣", "1f923"},
	{"⚡️", "26a1"},
	{"🍌", "1f34c"},
	{"🏆", "1f3c6"},
	{"💔", "1f494"},
	{"🤨", "1f928"},
	{"😐", "1f610"},
	{"🍓", "1f353"},
	{"🍾", "1f37e"},
	{"💋", "1f48b"},
	{"🖕", "1f595"},
	{"😈", "1f608"},
	{"😴", "1f634"},
	{"😭", "1f62d"},
	{"🤓", "1f913"},
	{"👻", "1f47b"},
	{"👨‍💻", "1f468-200d-1f4bb"},
	{"👀", "1f440"},
	{"🎃", "1f383"},
	{"🙈", "1f648"},
	{"😇", "1f607"},
	{"😨", "1f628"},
	{"🤝", "1f91d"},
	{"✍️", "270d"},
	{"🤗", "1f917"},
	{"🫡", "1fae1"},
	{"🎅", "1f385"},
	{"🎄", "1f384"},
	{"☃️", "2603"},
	{"💅", "1f485"},
	{"🤪", "1f92a"},
	{"🗿", "1f5ff"},
	{"🆒", "1f192"},
	{"💘", "1f498"},
	{"🙉", "1f649"},
	{"🦄", "1f984"},
	{"😘", "1f618"},
	{"💊", "1f48a"},
	{"🙊", "1f64a"},
	{"😎", "1f60e"},
	{"👾", "1f47e"},
	{"🤷‍♂️", "1f937-200d-2642"},
	{"🤷", "1f937"},
	{"🤷‍♀️", "1f937-200d-2640"},
	{"😡", "1f621"},
}

func buildEmojiDocument(id int64, def emojiDef) (*mtproto.Document, error) {
	var rawPNG []byte
	for _, baseDir := range []string{
		filepath.Join("teamgramd", "emoji", "64x64"),
		filepath.Join("emoji", "64x64"),
	} {
		fullPath := filepath.Join(baseDir, def.Codepoint+".png")
		data, err := os.ReadFile(fullPath)
		if err == nil {
			rawPNG = data
			break
		}
	}
	if rawPNG == nil {
		return nil, os.ErrNotExist
	}

	return mtproto.MakeTLDocument(&mtproto.Document{
		Id:            id,
		AccessHash:    id ^ 0x6D6F6A69,
		Date:          1,
		MimeType:      "image/png",
		Size2_INT64:   int64(len(rawPNG)),
		DcId:          1,
		FileReference: []byte{},
		Thumbs: []*mtproto.PhotoSize{
			mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
				Type:  "i",
				Bytes: rawPNG,
			}).To_PhotoSize(),
		},
		Attributes: []*mtproto.DocumentAttribute{
			mtproto.MakeTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
				W: 64,
				H: 64,
			}).To_DocumentAttribute(),
			mtproto.MakeTLDocumentAttributeSticker(&mtproto.DocumentAttribute{
				Alt:        def.Codepoint,
				Stickerset: &mtproto.InputStickerSet{},
				MaskCoords: &mtproto.MaskCoords{},
				Mask:       false,
			}).To_DocumentAttribute(),
		},
	}).To_Document(), nil
}
