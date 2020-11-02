package server

import (
	"bytes"

	"../../utils"
)

/**
Frame 2915: 55 bytes on wire (440 bits), 55 bytes captured (440 bits)
Null/Loopback
Internet Protocol Version 4, Src: 127.0.0.1, Dst: 127.0.0.1
Transmission Control Protocol, Src Port: 3306, Dst Port: 4447, Seq: 26119, Ack: 3586, Len: 11
MySQL Protocol
    Packet Length: 7
    Packet Number: 1
    Response Code: OK Packet (0x00)
    Affected Rows: 1
    Last INSERT ID: 1
    Server Status: 0x0002
        .... .... .... ...0 = In transaction: Not set
        .... .... .... ..1. = AUTO_COMMIT: Set
        .... .... .... .0.. = Multi query / Unused: Not set
        .... .... .... 0... = More results: Not set
        .... .... ...0 .... = Bad index used: Not set
        .... .... ..0. .... = No index used: Not set
        .... .... .0.. .... = Cursor exists: Not set
        .... .... 0... .... = Last row sent: Not set
        .... ...0 .... .... = Database dropped: Not set
        .... ..0. .... .... = No backslash escapes: Not set
        .... .0.. .... .... = Metadata changed: Not set
        .... 0... .... .... = Query was slow: Not set
        ...0 .... .... .... = PS Out Params: Not set
        ..0. .... .... .... = In Trans Readonly: Not set
        .0.. .... .... .... = Session state changed: Not set
    Warnings: 0

*/

type ResponseOK struct {
	ResponseCode uint8
	AffectedRows uint8
	LastInsertID int8
	ServerStatus uint16
	Warnings uint16
}

func (packet *ResponseOK) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	utils.WriteInteger(buffer, 1, uint64(packet.ResponseCode))
	utils.WriteInteger(buffer, 1, uint64(packet.AffectedRows))
	utils.WriteInteger(buffer, 1, uint64(packet.LastInsertID))
	utils.WriteInteger(buffer, 2, uint64(packet.ServerStatus))
	utils.WriteInteger(buffer, 2, uint64(packet.Warnings))
	return buffer.Bytes()
}

func (packet *ResponseOK) Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.ResponseCode = 	uint8(utils.ReadInteger(buffer, 1))
	packet.AffectedRows = 	uint8(utils.ReadInteger(buffer, 1))
	packet.LastInsertID = 	int8(utils.ReadInteger(buffer, 1))
	packet.ServerStatus = 	uint16(utils.ReadInteger(buffer, 2))
	packet.Warnings = 		uint16(utils.ReadInteger(buffer, 2))
}

