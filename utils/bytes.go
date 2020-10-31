package utils

import (
	"bytes"
	"encoding/binary"
)

/**
NOTE: MySQL use LittleEndian !!!
 */

/**
read bytes by binary.LittleEndian
 */
func ByteSliceToInt(bs []byte) (int64, error) {
	myBs := make([]byte, 64 / 8)
	copy(myBs, bs)
	buffer := bytes.NewBuffer(myBs)
	var num int64
	err := binary.Read(buffer, binary.LittleEndian, &num)
	return num, err
}

/**
write bytes by binary.LittleEndian
slice size is 8 (int64, 64 / 8 == 8)
if you want to the part of slice, please slice it by yourself
*/
func IntToByteSlice(num int64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, num)
	return bytesBuffer.Bytes(), err
}
