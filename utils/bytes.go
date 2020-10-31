package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

/**
NOTE: MySQL use LittleEndian !!!
 */

/**
read bytes by binary.LittleEndian
 */
func ByteSliceToInt(byteSlice []byte) int64 {
	myByteSlice := make([]byte, 64 / 8)
	copy(myByteSlice, byteSlice)
	buffer := bytes.NewBuffer(myByteSlice)
	var num int64
	err := binary.Read(buffer, binary.LittleEndian, &num)
	if err != nil {
		log.Printf("ByteSliceToInt err: [%+v], byteSlice: [%+v]", err, byteSlice)
	}
	return num
}

/**
write bytes by binary.LittleEndian
slice size is 8 (int64, 64 / 8 == 8)
if you want to the part of slice, please slice it by yourself
*/
func IntToByteSlice(num int64) []byte {
	buffer := bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, num)
	if err != nil {
		log.Printf("IntToByteSlice err: [%+v], byteSlice: [%+v]", err, num)
	}
	return buffer.Bytes()
}
