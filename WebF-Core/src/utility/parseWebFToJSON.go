package utility

//ParseWebFToJSON transforms a .webF into a .json file
func ParseWebFToJSON(input string, output string) {
	WriteJSON(WebFToJSON(ReadWebFFromFile(input)), output)
}

//ParseWebFToJSONFromString transforms a .WebF data into a .json file
func ParseWebFToJSONFromString(input string, output string) {
	WriteJSON(WebFToJSON(ReadWebF([]byte(input))), output)
}
