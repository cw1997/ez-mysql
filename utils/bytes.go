package utils

import (
	"encoding/binary"
)

/**
NOTE: MySQL use LittleEndian !!!
 */

/**
read bytes by binary.LittleEndian
 */
func ByteSliceToInt(byteSlice []byte) uint64 {
	myByteSlice := make([]byte, 64 / 8)
	copy(myByteSlice, byteSlice)
	return binary.LittleEndian.Uint64(myByteSlice)
}

/**
write bytes by binary.LittleEndian
slice size is 8 (int64, 64 / 8 == 8)
if you want to the part of slice, please slice it by yourself
*/
func IntToByteSlice(num uint64) []byte {
	byteSlice := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteSlice, num)
	return byteSlice
}
