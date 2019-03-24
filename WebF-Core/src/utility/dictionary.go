package utility

import (
	"strings"
)

//go:generate go run scripts/includeDictionary.go
// ^^^
// |||
// Execute this script before you can use the Dictionary function to get the list of words contained in the "src/dictionary /" folder
// and insert them into a file which will then be compiled

//Dictionary is a function that contains a list of words from different languages to free up space on .webf files
func Dictionary() (words []Words, length uint32) {
	var data []byte

	data = append(data, []byte(DicoBinfr)...)
	data = append(data, []byte(DicoBinen)...)

	wordList := strings.Split(string(data), "\r\n")

	for _, word := range wordList {
		words = append(words, Words{
			Length: uint8(len(word)),
			Word:   []byte(word),
		})
	}

	length = uint32(len(wordList))

	return
}
