package server

import (
	"bytes"

	"../../utils"
)

/**
Frame 38: 126 bytes on wire (1008 bits), 126 bytes captured (1008 bits)
Null/Loopback
Internet Protocol Version 4, Src: 127.0.0.1, Dst: 127.0.0.1
Transmission Control Protocol, Src Port: 3306, Dst Port: 4447, Seq: 1, Ack: 1, Len: 82
MySQL Protocol
    Packet Length: 78
    Packet Number: 0
    Server Greeting
        Protocol: 10
        Version: 5.7.30-log
        Thread ID: 11
        Salt: lr;&\025?K*
        Server Capabilities: 0xffff
            .... .... .... ...1 = Long Password: Set
            .... .... .... ..1. = Found Rows: Set
            .... .... .... .1.. = Long Column Flags: Set
            .... .... .... 1... = Connect With Database: Set
            .... .... ...1 .... = Don't Allow database.table.column: Set
            .... .... ..1. .... = Can use compression protocol: Set
            .... .... .1.. .... = ODBC Client: Set
            .... .... 1... .... = Can Use LOAD DATA LOCAL: Set
            .... ...1 .... .... = Ignore Spaces before '(': Set
            .... ..1. .... .... = Speaks 4.1 protocol (new flag): Set
            .... .1.. .... .... = Interactive Client: Set
            .... 1... .... .... = Switch to SSL after handshake: Set
            ...1 .... .... .... = Ignore sigpipes: Set
            ..1. .... .... .... = Knows about transactions: Set
            .1.. .... .... .... = Speaks 4.1 protocol (old flag): Set
            1... .... .... .... = Can do 4.1 authentication: Set
        Server Language: latin1 COLLATE latin1_swedish_ci (8)
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
        Extended Server Capabilities: 0xc1ff
            .... .... .... ...1 = Multiple statements: Set
            .... .... .... ..1. = Multiple results: Set
            .... .... .... .1.. = PS Multiple results: Set
            .... .... .... 1... = Plugin Auth: Set
            .... .... ...1 .... = Connect attrs: Set
            .... .... ..1. .... = Plugin Auth LENENC Client Data: Set
            .... .... .1.. .... = Client can handle expired passwords: Set
            .... .... 1... .... = Session variable tracking: Set
            .... ...1 .... .... = Deprecate EOF: Set
            1100 000. .... .... = Unused: 0x60
        Authentication Plugin Length: 21
        Unused: 00000000000000000000
        Salt: P@g\030lH*Orn,\021
        Authentication Plugin: mysql_native_password

*/

type Greeting struct {
	Protocol             uint8
	Version              string
	ThreadId             uint32
	Salt                 []byte
	// padding 1 byte : 0x00
	ServerCapabilities   uint16
	ServerLanguage       uint8
	ServerStatus         uint16
	ExtendedServerStatus uint16
	ExtendedSaltLength   uint8
	Reverse              []byte
	// ExtendedSalt length is (ExtendedSaltLength + 8)
	ExtendedSalt               []byte
	AuthenticationPlugin       string
}

func (packet *Greeting) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	utils.WriteInteger(buffer, 1, uint64(packet.Protocol))
	utils.WriteNullTerminatedString(buffer, packet.Version)
	utils.WriteInteger(buffer, 4, uint64(packet.ThreadId))
	utils.WriteNullTerminatedString(buffer, string(packet.Salt))
	utils.WriteInteger(buffer, 2, uint64(packet.ServerCapabilities))
	utils.WriteInteger(buffer, 1, uint64(packet.ServerLanguage))
	utils.WriteInteger(buffer, 2, uint64(packet.ServerStatus))
	utils.WriteInteger(buffer, 2, uint64(packet.ExtendedServerStatus))
	utils.WriteInteger(buffer, 1, uint64(packet.ExtendedSaltLength))
	utils.WriteRepeat(buffer, []byte{0}, 10)
	utils.WriteNullTerminatedString(buffer, string(packet.ExtendedSalt))
	utils.WriteNullTerminatedString(buffer, packet.AuthenticationPlugin)
	return buffer.Bytes()
}

func (packet *Greeting) Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.Protocol = uint8(utils.ReadInteger(buffer, 1))
	packet.Version = utils.ReadNullTerminatedString(buffer)
	packet.ThreadId = uint32(utils.ReadInteger(buffer, 4))
	packet.Salt = []byte(utils.ReadNullTerminatedString(buffer))
	packet.ServerCapabilities = uint16(utils.ReadInteger(buffer, 2))
	packet.ServerLanguage = uint8(utils.ReadInteger(buffer, 1))
	packet.ServerStatus = uint16(utils.ReadInteger(buffer, 2))
	packet.ExtendedServerStatus = uint16(utils.ReadInteger(buffer, 2))
	packet.ExtendedSaltLength = uint8(utils.ReadInteger(buffer, 1))
	packet.Reverse = buffer.Next(10)
	// packet.ExtendedSaltLength - 8 - 1, skip the end of 0
	packet.ExtendedSalt = buffer.Next(int(packet.ExtendedSaltLength - 8 - 1))
	//skip the end of 0
	buffer.Next(1)
	packet.AuthenticationPlugin = utils.ReadNullTerminatedString(buffer)
}
