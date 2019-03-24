package utility

//ParseJSONToWebF transforms a json file into a .WebF file
func ParseJSONToWebF(input string, output string) {
	WriteWebF(JSONToWebF(ReadJSONFromFile(input)), output)
}

//ParseJSONToWebFFromString transforms a json string into a .WebF file
func ParseJSONToWebFFromString(input string, output string) {
	WriteWebF(JSONToWebF(ReadJSON([]byte(input))), output)
}
