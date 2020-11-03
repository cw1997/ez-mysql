package client

import (
	"../../utils"
	"bytes"
)

/**
Frame 40: 226 bytes on wire (1808 bits), 226 bytes captured (1808 bits)
Null/Loopback
Internet Protocol Version 4, Src: 127.0.0.1, Dst: 127.0.0.1
Transmission Control Protocol, Src Port: 4447, Dst Port: 3306, Seq: 1, Ack: 83, Len: 182
MySQL Protocol
    Packet Length: 178
    Packet Number: 1
    Login Request
        Client Capabilities: 0xa685
            .... .... .... ...1 = Long Password: Set
            .... .... .... ..0. = Found Rows: Not set
            .... .... .... .1.. = Long Column Flags: Set
            .... .... .... 0... = Connect With Database: Not set
            .... .... ...0 .... = Don't Allow database.table.column: Not set
            .... .... ..0. .... = Can use compression protocol: Not set
            .... .... .0.. .... = ODBC Client: Not set
            .... .... 1... .... = Can Use LOAD DATA LOCAL: Set
            .... ...0 .... .... = Ignore Spaces before '(': Not set
            .... ..1. .... .... = Speaks 4.1 protocol (new flag): Set
            .... .1.. .... .... = Interactive Client: Set
            .... 0... .... .... = Switch to SSL after handshake: Not set
            ...0 .... .... .... = Ignore sigpipes: Not set
            ..1. .... .... .... = Knows about transactions: Set
            .0.. .... .... .... = Speaks 4.1 protocol (old flag): Not set
            1... .... .... .... = Can do 4.1 authentication: Set
        Extended Client Capabilities: 0x007f
            .... .... .... ...1 = Multiple statements: Set
            .... .... .... ..1. = Multiple results: Set
            .... .... .... .1.. = PS Multiple results: Set
            .... .... .... 1... = Plugin Auth: Set
            .... .... ...1 .... = Connect attrs: Set
            .... .... ..1. .... = Plugin Auth LENENC Client Data: Set
            .... .... .1.. .... = Client can handle expired passwords: Set
            .... .... 0... .... = Session variable tracking: Not set
            .... ...0 .... .... = Deprecate EOF: Not set
            0000 000. .... .... = Unused: 0x00
        MAX Packet: 1073741824
        Charset: utf8 COLLATE utf8_general_ci (33)
        Unused: 0000000000000000000000000000000000000000000000
        Username: root
        Password: 8dfb928ea879596be81dfee465a4c2e7ad234043
        Client Auth Plugin: mysql_native_password
        Connection Attributes
            Connection Attributes length: 97
            Connection Attribute - _os: Win64
                Connection Attribute Name Length: 3
                Connection Attribute Name: _os
                Connection Attribute Name Length: 5
                Connection Attribute Value: Win64
            Connection Attribute - _client_name: libmysql
                Connection Attribute Name Length: 12
                Connection Attribute Name: _client_name
                Connection Attribute Name Length: 8
                Connection Attribute Value: libmysql
            Connection Attribute - _pid: 25672
                Connection Attribute Name Length: 4
                Connection Attribute Name: _pid
                Connection Attribute Name Length: 5
                Connection Attribute Value: 25672
            Connection Attribute - _thread: 16740
                Connection Attribute Name Length: 7
                Connection Attribute Name: _thread
                Connection Attribute Name Length: 5
                Connection Attribute Value: 16740
            Connection Attribute - _platform: AMD64
                Connection Attribute Name Length: 9
                Connection Attribute Name: _platform
                Connection Attribute Name Length: 5
                Connection Attribute Value: AMD64
            Connection Attribute - _client_version: 10.1.24
                Connection Attribute Name Length: 15
                Connection Attribute Name: _client_version
                Connection Attribute Name Length: 7
                Connection Attribute Value: 10.1.24

*/

type Login struct {
	ClientCapabilities         uint16
	ExtendedClientCapabilities uint16
	MAXPacket                  uint32
	Charset                    uint8
	//Unused 23 bytes
	Username             string
	PasswordLength       uint8 // 20 bytes
	Password             []byte // 20 bytes
	ClientAuthPlugin     string
	ConnectionAttributes []byte
}

func (packet *Login) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	utils.WriteInteger(buffer, 2, uint64(packet.ClientCapabilities))
	utils.WriteInteger(buffer, 2, uint64(packet.ExtendedClientCapabilities))
	utils.WriteInteger(buffer, 4, uint64(packet.MAXPacket))
	utils.WriteInteger(buffer, 1, uint64(packet.Charset))
	utils.WriteRepeat(buffer, []byte{0}, 23)
	utils.WriteNullTerminatedString(buffer, packet.Username)
	utils.WriteLengthCodedBinary(buffer, packet.Password)
	utils.WriteNullTerminatedString(buffer, packet.ClientAuthPlugin)
	utils.WriteLengthCodedBinary(buffer, packet.ConnectionAttributes)
	return buffer.Bytes()
}

func (packet *Login) Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.ClientCapabilities = uint16(utils.ReadInteger(buffer, 2))
	packet.ExtendedClientCapabilities = uint16(utils.ReadInteger(buffer, 2))
	packet.MAXPacket = uint32(utils.ReadInteger(buffer, 4))
	packet.Charset = uint8(utils.ReadInteger(buffer, 1))
	buffer.Next(23)
	packet.Username = utils.ReadNullTerminatedString(buffer)
	password, passwordLength := utils.ReadLengthCodedBinary(buffer)
	packet.Password, packet.PasswordLength = password, uint8(passwordLength)
	packet.ClientAuthPlugin = utils.ReadNullTerminatedString(buffer)
	packet.ConnectionAttributes, _ = utils.ReadLengthCodedBinary(buffer)
}

/**
reference: https://dev.mysql.com/doc/internals/en/secure-password-authentication.html
SHA1( password ) XOR SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
 */
func MysqlNativePassword(scramble []byte, password string) []byte {
	if len(scramble) != 20 {
		panic("scramble length != 20")
	}
	passwordHashed := utils.SHA1([]byte(password))
	//log.Println(scramble, len(scramble), passwordHashed, len(passwordHashed))
	tmp := utils.SHA1(append(scramble, utils.SHA1(passwordHashed)...))
	for i := range tmp {
		tmp[i] ^= passwordHashed[i]
	}
	return tmp
}
