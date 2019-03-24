package utility

import (
	"bytes"
	"encoding/binary"
)

//ReadUint16 transforms an array of bytes into uint16 type
func ReadUint16(data []byte, index int) uint16 {
	tmp := []byte{data[index], data[index+1]}
	var tmp16 uint16
	buf := bytes.NewBuffer(tmp)
	binary.Read(buf, binary.LittleEndian, &tmp16)

	return tmp16
}
