package cli

import (
	"fmt"
	"os"

	"./WebF-Core/src/utility"
)

func main() {
	// WebF official Github repository : https://github.com/ThomasT404/WebF
	// WebF-CLI official Github repository : https://github.com/ThomasT404/WebF-CLI
	// WebF Version : 0.1 (WIP)
	// WebF-CLI Version : 0.1 (WIP)
	// Author : Thomas Tastet

	argsWithoutProg := os.Args[1:]
	lenn := len(argsWithoutProg)

	if lenn == 0 {
		utility.Error("You have not given any settings")
	} else {
		if lenn == 2 {
			if argsWithoutProg[0] == "parseFile" {
				fmt.Println(utility.JSONToWebF(utility.ReadJSONFromFile(argsWithoutProg[1])))
			} else if argsWithoutProg[0] == "parse" {
				fmt.Println(utility.JSONToWebF(utility.ReadJSON([]byte(argsWithoutProg[1]))))
			} else if argsWithoutProg[0] == "decodeFile" {
				fmt.Println(utility.WebFToJSON(utility.ReadWebFFromFile(argsWithoutProg[1])))
			} else if argsWithoutProg[0] == "decode" {
				fmt.Println(utility.WebFToJSON(utility.ReadWebF([]byte(argsWithoutProg[1]))))
			}
		} else if lenn == 3 {
			if argsWithoutProg[0] == "parseFile" {
				utility.ParseJSONToWebF(argsWithoutProg[1], argsWithoutProg[2])
			} else if argsWithoutProg[0] == "parse" {
				utility.ParseJSONToWebFFromString(argsWithoutProg[1], argsWithoutProg[2])
			} else if argsWithoutProg[0] == "decodeFile" {
				utility.ParseWebFToJSON(argsWithoutProg[1], argsWithoutProg[2])
			} else if argsWithoutProg[0] == "decode" {
				utility.ParseWebFToJSONFromString(argsWithoutProg[1], argsWithoutProg[2])
			}
		} else {
			utility.Error("2 or 3 parameters expected")
		}
	}
}
