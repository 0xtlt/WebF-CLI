package utility

import "io/ioutil"

//ReadWebFFromFile reads a .webf file and retransmits the data in a WebF variable
func ReadWebFFromFile(path string) (content WebF) {
	data, err := ioutil.ReadFile(path)
	Check(err)

	content = ReadWebF(data)
	return
}

//ReadWebFFromBytes
