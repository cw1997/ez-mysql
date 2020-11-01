package server

import (
	"bytes"

	"../../utils"
)

/**
Frame 4156: 65 bytes on wire (520 bits), 65 bytes captured (520 bits)
Null/Loopback
Internet Protocol Version 4, Src: 127.0.0.1, Dst: 127.0.0.1
Transmission Control Protocol, Src Port: 4447, Dst Port: 3306, Seq: 4351, Ack: 32002, Len: 21
MySQL Protocol
    Packet Length: 17
    Packet Number: 0
    Request Command Query
        Command: Query (3)
        Statement: select version()

*/

type Request struct {
	Command uint8
	Statement string
}

func (packet *Request) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(utils.IntToByteSlice(uint64(packet.Command))[0:1])
	buffer.WriteString(packet.Statement)
	return buffer.Bytes()
}

func (packet *Request) Resolve(byteSlice []byte, length uint32) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.Command = uint8(utils.ByteSliceToInt(buffer.Next(1)))
	packet.Statement = string(buffer.Next(int(length - 1)))
}
