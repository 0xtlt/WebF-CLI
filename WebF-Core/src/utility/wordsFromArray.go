package utility

import "strings"

//WordsFromString returns an array of words containing each word to be added to the dictionary of the .webf file
func WordsFromString(text string) (words []Words) {
	tmp := strings.Split(text, " ")

	dico, _ := Dictionary()

	for _, word := range tmp {
		if IndexWord(words, word) == -1 && IndexWord(dico, word) == -1 {
			words = append(words, Words{
				Length: uint8(len([]byte(word))),
				Word:   []byte(word),
			})
		}
	}

	return
}
