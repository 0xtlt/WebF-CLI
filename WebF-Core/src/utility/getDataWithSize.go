package utility

//GetDataWithSize first reads the first byte to determine what type is the information that follows and then adjusts the following information and returns it as uint32
func GetDataWithSize(data []byte, cursor *int) (content uint32) {
	switch data[*cursor] {
	case 0x00:
		content = uint32(data[*cursor+1])
		*cursor += 2
		break
	case 0x01:
		content = uint32(ReadUint16(data, *cursor+1))
		*cursor += 3
		break
	case 0x02:
		content = ReadUint32(data, *cursor+1)
		*cursor += 5
		break
	}
	return
}
