package utility

import (
	"bytes"
	"encoding/binary"
)

//ReadUint32 transforms an array of bytes into uint32 type
func ReadUint32(data []byte, index int) uint32 {
	tmp := []byte{data[index], data[index+1], data[index+2], data[index+3]}
	var tmp32 uint32
	buf := bytes.NewBuffer(tmp)
	binary.Read(buf, binary.LittleEndian, &tmp32)

	return tmp32
}
