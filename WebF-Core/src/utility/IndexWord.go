package utility

//IndexWord searches and returns the position of the word in the word list.
func IndexWord(words []Words, search string) (re int) {
	re = -1

	for index, x := range words {
		if string(x.Word) == search {
			re = index
		}
	}

	return
}
