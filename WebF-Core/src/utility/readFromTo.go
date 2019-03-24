package utility

//ReadFromTo returns an array of bytes containing the bytes that is between the given index and the given length
func ReadFromTo(data []byte, from int, length int) (c []byte) {
	for n := from; n < from+length; n++ {
		c = append(c, data[n])
	}
	return
}
