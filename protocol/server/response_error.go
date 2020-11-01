package server

import (
	"bytes"

	"../../utils"
)

/**
FFrame 4100: 97 bytes on wire (776 bits), 97 bytes captured (776 bits)
Null/Loopback
Internet Protocol Version 4, Src: 127.0.0.1, Dst: 127.0.0.1
Transmission Control Protocol, Src Port: 3306, Dst Port: 4447, Seq: 31949, Ack: 4351, Len: 53
MySQL Protocol
    Packet Length: 49
    Packet Number: 1
    Response Code: ERR Packet (0xff)
    Error Code: 1054
    SQL state: 42S22
    Error message: Unknown column 'version' in 'field list'

*/

type ResponseError struct {
	ResponseCode uint8
	ErrorCode    uint16
	SQLState     []byte // 5 bytes
	Errormessage string
}

func (packet *ResponseError) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(utils.IntToByteSlice(uint64(packet.ResponseCode))[0:1])
	buffer.Write(utils.IntToByteSlice(uint64(packet.ErrorCode))[0:2])
	buffer.Write([]byte("#",))
	buffer.Write(packet.SQLState)
	buffer.WriteString(packet.Errormessage)
	return buffer.Bytes()
}

func (packet *ResponseError) Resolve(byteSlice []byte, length uint32) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.ResponseCode = uint8(utils.ByteSliceToInt(buffer.Next(8 / 8)))
	packet.ErrorCode = uint16(utils.ByteSliceToInt(buffer.Next(16 / 8)))
	// skip '#'
	buffer.Next(1)
	packet.SQLState = buffer.Next(5)
	packet.Errormessage = string(buffer.Next(int(length - 1 - 1 - 2 - 5)))
}
