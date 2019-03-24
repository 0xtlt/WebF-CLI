package utility

//JSONToWebF adapts a JsonWebF type in WebF type
func JSONToWebF(json JSONWebF) (content WebF) {
	title := json.Title
	var metas []Meta

	for _, meta := range json.Meta {
		metas = append(metas, Meta{
			Keylength:   uint8(len(meta.Key)),
			Key:         meta.Key,
			Valuelength: uint16(len(meta.Value)),
			Value:       meta.Value,
		})
	}

	text := json.Content

	header := Header{
		Filetype:      uint16(0xFF), //it will change later
		Titlelength:   uint16(len(title)),
		Metalength:    uint16(len(metas)),
		Wordslength:   0,
		Contentlength: 0,
	}

	content = WebF{
		Header:  header,
		Title:   title,
		Meta:    metas,
		Words:   WordsFromString(string(text)),
		Content: nil,
	}

	content.Content = WordsPosition(content.Words, string(text))

	content.Header.Wordslength = uint32(len(content.Words))
	content.Header.Contentlength = uint64(len(content.Content))

	return
}
