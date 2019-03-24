package utility

//MixPosWithWords returns the fusion of the Words array and positions array
func MixPosWithWords(words []Words, content []uint32) (text string) {
	dico, _ := Dictionary()

	for _, position := range content {
		if (position & 0x80000000) == 0x80000000 {
			text += string(dico[int(position&0x7FFFFFFF)].Word) + " "
		} else if len(words) > 0 {
			text += string(words[int(position)].Word) + " "
		}
	}

	return
}
