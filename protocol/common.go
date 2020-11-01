package protocol

import (
	"../utils"
	"bytes"
)

const SERVER_VERSION = "5.7.26-log"

type MySQLMessage struct {
	PayloadLength uint32
	SequenceId    uint8
	Payload       []byte
}

func (message *MySQLMessage)Build() []byte {
	//size = 32 / 8 + 8 / 8 + message.PayloadLength
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(utils.IntToByteSlice(uint64(message.PayloadLength))[0:3])
	buffer.Write(utils.IntToByteSlice(uint64(message.SequenceId))[0:1])
	buffer.Write(message.Payload)
	return buffer.Bytes()
}

func (message *MySQLMessage)Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	message.PayloadLength = uint32(utils.ByteSliceToInt(buffer.Next(3)))
	message.SequenceId = uint8(utils.ByteSliceToInt(buffer.Next(1)))
	message.Payload = buffer.Next(int(message.PayloadLength))
}
