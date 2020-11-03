package protocol

import (
	"../utils"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
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

func ReadMySQLMessage(conn net.Conn) (header *Header, payload []byte, err error) {
	buffer := make([]byte, 4)
	n, err := io.ReadFull(conn, buffer)
	if err != nil {
		log.Println("io.ReadFull(conn, buffer) read header: ", n, err)
		return
	}

	header = new(Header)
	header.Resolve(buffer)
	payloadLength := header.PayloadLength

	fmt.Printf("ReadMySQLMessage\t-> header: %+v \n", header)

	payload = make([]byte, payloadLength)
	n, err = io.ReadFull(conn, payload)
	if err != nil {
		log.Println("io.ReadFull(conn, buffer) read payload: ", n, err)
		return
	}
	return
}

func WriteMySQLMessage(conn net.Conn, payload []byte, sequenceId uint8) (n int, err error) {
	header := new(Header)
	header.PayloadLength = uint32(len(payload))
	header.SequenceId = sequenceId

	fmt.Printf("WriteMySQLMessage\t-> header: %+v \n", header)

	n, err = conn.Write(header.Build())
	if err != nil {
		return
	}
	n, err = conn.Write(payload)
	if err != nil {
		return
	}
	return
}

