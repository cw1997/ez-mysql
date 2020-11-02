package protocol

import (
	"../utils"
	"bytes"
)

const (
	payloadLengthSize = 3
	sequenceIdSize    = 1
)


type Header struct {
	PayloadLength uint32
	SequenceId    uint8
}

func (packet *Header)Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	utils.WriteInteger(buffer, payloadLengthSize, uint64(packet.PayloadLength))
	utils.WriteInteger(buffer, sequenceIdSize, uint64(packet.SequenceId))
	return buffer.Bytes()
}

func (packet *Header)Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.PayloadLength = uint32(utils.ReadInteger(buffer, payloadLengthSize))
	packet.SequenceId = uint8(utils.ReadInteger(buffer, sequenceIdSize))
}


type MySQLMessage struct {
	Header Header
	Payload []byte
}

func (packet *MySQLMessage)Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	utils.WriteInteger(buffer, payloadLengthSize, uint64(packet.Header.PayloadLength))
	utils.WriteInteger(buffer, sequenceIdSize, uint64(packet.Header.SequenceId))
	buffer.Write(packet.Payload)
	return buffer.Bytes()
}

func (packet *MySQLMessage)Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.Header.PayloadLength = uint32(utils.ReadInteger(buffer, payloadLengthSize))
	packet.Header.SequenceId = uint8(utils.ReadInteger(buffer, sequenceIdSize))
	packet.Payload = buffer.Next(int(packet.Header.PayloadLength))
}
