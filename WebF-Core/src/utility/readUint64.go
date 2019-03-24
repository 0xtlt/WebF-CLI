package utility

import (
	"bytes"
	"encoding/binary"
)

//ReadUint64 transforms an array of bytes into uint64 type
func ReadUint64(data []byte, index int) uint64 {
	tmp := []byte{data[index], data[index+1], data[index+2], data[index+3], data[index+4], data[index+5], data[index+6], data[index+7]}
	var tmp64 uint64
	buf := bytes.NewBuffer(tmp)
	binary.Read(buf, binary.LittleEndian, &tmp64)

	return tmp64
}
