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

func ReadInteger(buffer *bytes.Buffer, length uint8) int64 {
	number := buffer.Next(int(length))
	return int64(ByteSliceToInt(number))
}

func WriteInteger(buffer *bytes.Buffer, length uint8, number uint64) {
	buffer.Write(IntToByteSlice(number)[:length])
}


func ReadNullTerminatedString(buffer *bytes.Buffer) string {
	str, _ := buffer.ReadString(0)
	return str[:len(str)-1]
}

func WriteNullTerminatedString(buffer *bytes.Buffer, str string) {
	buffer.WriteString(str)
	buffer.WriteByte(0)
}


func ReadLengthCodedBinary(buffer *bytes.Buffer) (byteSlice []byte, length uint64) {
	firstByte, _ := buffer.ReadByte()
	if firstByte < 251 {
		length = uint64(firstByte)
		byteSlice = buffer.Next(int(firstByte))
	} else {
		switch firstByte {
		case 252:
			length = ByteSliceToInt(buffer.Next(2))
			byteSlice = buffer.Next(int(length))
			break
		case 253:
			length = ByteSliceToInt(buffer.Next(3))
			byteSlice = buffer.Next(int(length))
			break
		case 254:
			length = ByteSliceToInt(buffer.Next(8))
			byteSlice = buffer.Next(int(length))
			break
		default:
			length = 0
			byteSlice = []byte{}
		}
	}
	return byteSlice, length
}

func WriteLengthCodedBinary(buffer *bytes.Buffer, byteSlice []byte) {
	length := len(byteSlice)
	if length < 251 {
		buffer.WriteByte(byte(length))
	} else {
		if length >= 2 << (3 * 8) {
			buffer.WriteByte(254)
			WriteInteger(buffer, 8, uint64(length))
		} else if length >= 2 << (2 * 8) {
			buffer.WriteByte(253)
			WriteInteger(buffer, 3, uint64(length))
		} else {
			buffer.WriteByte(252)
			WriteInteger(buffer, 2, uint64(length))
		}
	}
	buffer.Write(byteSlice)
}
