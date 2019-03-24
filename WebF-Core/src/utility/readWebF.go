package utility

//ReadWebF reads a .webf file and retransmits the data in a WebF variable
func ReadWebF(data []byte) (content WebF) {
	var cursor int

	content.Header.Filetype = ReadUint16(data, cursor)
	cursor += 2

	content.Header.Titlelength = ReadUint16(data, cursor)
	cursor += 2

	content.Header.Metalength = ReadUint16(data, cursor)
	cursor += 2

	content.Header.Wordslength = ReadUint32(data, cursor)
	cursor += 4

	content.Header.Contentlength = ReadUint64(data, cursor)
	cursor += 8

	//Meta
	for n := 0; n < int(content.Header.Metalength); n++ {
		var tmpMeta Meta

		tmpMeta.Keylength = data[cursor]
		cursor++

		tmpMeta.Key = string(ReadFromTo(data, cursor, int(tmpMeta.Keylength)))
		cursor += int(tmpMeta.Keylength)

		tmpMeta.Valuelength = ReadUint16(data, cursor)
		cursor += 2

		tmpMeta.Value = string(ReadFromTo(data, cursor, int(tmpMeta.Valuelength)))
		cursor += int(tmpMeta.Valuelength)

		content.Meta = append(content.Meta, tmpMeta)
	}

	//Title
	content.Title = string(ReadFromTo(data, cursor, int(content.Header.Titlelength)))
	cursor += int(content.Header.Titlelength)

	// fmt.Println("curs start ", cursor)

	//Words
	for n := 0; n < int(content.Header.Wordslength); n++ {
		var tmpWord Words

		tmpWord.Length = data[cursor]
		cursor++

		tmpWord.Word = ReadFromTo(data, cursor, int(tmpWord.Length))
		cursor += int(tmpWord.Length)

		content.Words = append(content.Words, tmpWord)
	}

	//Content
	for n := 0; n < int(content.Header.Contentlength); n++ {
		tmpData := GetDataWithSize(data, &cursor)
		content.Content = append(content.Content, tmpData)
	}

	return
}
