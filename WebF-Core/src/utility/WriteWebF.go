package utility

import (
	"encoding/binary"
	"math"
	"os"
)

//WriteWebF saves WebF type data in a .webf file
func WriteWebF(content WebF, output string) {
	f, err := os.Create(output)
	Check(err)
	defer f.Close()

	//header
	//filetype
	a := make([]byte, 2)
	binary.LittleEndian.PutUint16(a, content.Header.Filetype)
	f.Write(a)

	//Titlelength
	a = make([]byte, 2)
	binary.LittleEndian.PutUint16(a, content.Header.Titlelength)
	f.Write(a)

	//Metalength
	a = make([]byte, 2)
	binary.LittleEndian.PutUint16(a, content.Header.Metalength)
	f.Write(a)

	//Wordslength
	a = make([]byte, 4)
	binary.LittleEndian.PutUint32(a, content.Header.Wordslength)
	f.Write(a)

	//Contentlength
	a = make([]byte, 8)
	binary.LittleEndian.PutUint64(a, content.Header.Contentlength)
	f.Write(a)

	//Meta
	for _, meta := range content.Meta {
		//Keylength
		f.Write([]byte{meta.Keylength})

		//Key
		f.Write([]byte(meta.Key))

		//Valuelength
		a = make([]byte, 2)
		binary.LittleEndian.PutUint16(a, meta.Valuelength)
		f.Write(a)

		//Value
		f.Write([]byte(meta.Value))
	}

	//Title
	f.Write([]byte(content.Title))

	//Words
	for _, word := range content.Words {
		//length
		f.Write([]byte{word.Length})
		//word
		f.Write(word.Word)
	}

	//Content
	//the first byte written is the type of the next content position
	//ex: if content position is < 255, the first byte is 0x00 and the position saved after that will be in uint8 bits
	for _, content := range content.Content {
		if int(content) < math.MaxUint8 {
			f.Write([]byte{uint8(0x00)})
			f.Write([]byte{uint8(content)})
		} else if int(content) < math.MaxUint16 {
			f.Write([]byte{uint8(0x01)})
			a = make([]byte, 2)

			binary.LittleEndian.PutUint32(a, content)
			f.Write(a)
		} else if int(content) < math.MaxUint32 {
			f.Write([]byte{uint8(0x02)})
			a = make([]byte, 4)

			binary.LittleEndian.PutUint32(a, content)
			f.Write(a)
		}
	}
}
