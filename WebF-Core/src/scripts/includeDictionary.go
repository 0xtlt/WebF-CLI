package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Reads all .txt files in the "src/doctionary" folder
// and encodes them as strings literals in src/utility/dictionaryBin.go
func main() {
	fs, _ := ioutil.ReadDir("./dictionary")
	out, _ := os.Create("./utility/dictionaryBin.go")
	out.Write([]byte("package utility \n\n//DicoBin\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".txt") {
			out.Write([]byte("DicoBin" + (strings.TrimSuffix(f.Name(), ".txt")[4:len(strings.TrimSuffix(f.Name(), ".txt"))]) + " = `"))
			f, _ := os.Open("./dictionary/" + f.Name())
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
