package utility

import (
	"os"
)

//WriteJSON function writes string that contains json data in json file
func WriteJSON(content string, output string) {
	f, err := os.Create(output)
	Check(err)
	defer f.Close()

	f.Write([]byte(content))
}
