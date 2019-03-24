package utility

import (
	"strings"
)

//WordsPosition returns a uint32 array containing the position of each word in the dictionaries
func WordsPosition(words []Words, text string) (position []uint32) {
	tmp := strings.Split(text, " ")

	dico, _ := Dictionary()

	for _, word := range tmp {
		dicoIndex := IndexWord(dico, word)

		if dicoIndex != -1 {
			pos := dicoIndex | 0x80000000
			position = append(position, uint32(pos))
			//7FFF
		} else {
			pos := IndexWord(words, word)
			if pos != -1 {
				position = append(position, uint32(pos))
			}
		}
	}

	return
}
