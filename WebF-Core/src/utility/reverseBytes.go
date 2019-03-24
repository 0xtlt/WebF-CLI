package utility

//ReverseBytes changes the order of the byte array
func ReverseBytes(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(ReverseBytes(input[1:]), input[0])
}
