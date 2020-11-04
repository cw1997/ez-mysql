package server

import (
	"../../utils"
	"bytes"
)

type ResponseResultSet struct {
	ResponseField   []ResponseField
	ResponseRowData []ResponseRowData
}

/**
Frame 104: 454 bytes on wire (3632 bits), 454 bytes captured (3632 bits) on interface unknown, id 0
Null/Loopback
Internet Protocol Version 4, Src: 127.0.0.1, Dst: 127.0.0.1
Transmission Control Protocol, Src Port: 3306, Dst Port: 4447, Seq: 26566, Ack: 3860, Len: 410
MySQL Protocol
MySQL Protocol
MySQL Protocol
    Packet Length: 50
    Packet Number: 3
    Catalog: def
    Database: test
    Table: user
    Original table: user
    Name: username
    Original name: username
    Charset number: utf8mb4 COLLATE utf8mb4_general_ci (45)
    Length: 1020
    Type: FIELD_TYPE_VAR_STRING (253)
    Flags: 0x1001
        .... .... .... ...1 = Not null: Set
        .... .... .... ..0. = Primary key: Not set
        .... .... .... .0.. = Unique key: Not set
        .... .... .... 0... = Multiple key: Not set
        .... .... ...0 .... = Blob: Not set
        .... .... ..0. .... = Unsigned: Not set
        .... .... .0.. .... = Zero fill: Not set
        .... .... 0... .... = Binary: Not set
        .... ...0 .... .... = Enum: Not set
        .... ..0. .... .... = Auto increment: Not set
        .... .0.. .... .... = Timestamp: Not set
        .... 0... .... .... = Set: Not set
    Decimals: 0
MySQL Protocol
MySQL Protocol
MySQL Protocol
MySQL Protocol
MySQL Protocol
MySQL Protocol
MySQL Protocol

 */

type ResponseField struct {
	Catalog       string
	Database      string
	Table         string
	OriginalTable string
	Name          string
	OriginalName  string
	CharsetNumber uint16
	Length        uint32
	Type          uint8
	Flags         uint16
	Decimals      uint8
}

func (packet *ResponseField) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})

	utils.WriteLengthCodedBinary(buffer, []byte(packet.Catalog))
	utils.WriteLengthCodedBinary(buffer, []byte(packet.Database))
	utils.WriteLengthCodedBinary(buffer, []byte(packet.Table))
	utils.WriteLengthCodedBinary(buffer, []byte(packet.OriginalTable))
	utils.WriteLengthCodedBinary(buffer, []byte(packet.Name))
	utils.WriteLengthCodedBinary(buffer, []byte(packet.OriginalName))
	buffer.WriteByte(0x0c)
	utils.WriteInteger(buffer, 2, uint64(packet.CharsetNumber))
	utils.WriteInteger(buffer, 4, uint64(packet.Length))
	utils.WriteInteger(buffer, 1, uint64(packet.Type))
	utils.WriteInteger(buffer, 2, uint64(packet.Flags))
	utils.WriteInteger(buffer, 1, uint64(packet.Decimals))
	utils.WriteRepeat(buffer, []byte{0}, 2)

	//typeOfPacket := reflect.TypeOf(*packet)
	//valueOfPacket := reflect.ValueOf(*packet)
	//numField := typeOfPacket.NumField()
	//for i := 0; i < numField; i++ {
	//	field := typeOfPacket.Field(i)
	//	fieldKind := field.Type.Kind()
	//	switch fieldKind {
	//	case reflect.String:
	//		value := valueOfPacket.Field(i).String()
	//		utils.WriteLengthCodedBinary(buffer, []byte(value))
	//		break
	//	case reflect.Uint8:
	//		value := valueOfPacket.Field(i).Uint()
	//		utils.WriteInteger(buffer, 1, value)
	//		break
	//	case reflect.Uint16:
	//		value := valueOfPacket.Field(i).Uint()
	//		utils.WriteInteger(buffer, 2, value)
	//		break
	//	case reflect.Uint32:
	//		value := valueOfPacket.Field(i).Uint()
	//		utils.WriteInteger(buffer, 4, value)
	//		break
	//	}
	//}

	return buffer.Bytes()
}

func (packet *ResponseField) Resolve(sliceByte []byte) {
	buffer := bytes.NewBuffer(sliceByte)

	Catalog, _ := utils.ReadLengthCodedBinary(buffer)
	packet.Catalog = string(Catalog)
	Database, _ := utils.ReadLengthCodedBinary(buffer)
	packet.Database = string(Database)
	Table, _ := utils.ReadLengthCodedBinary(buffer)
	packet.Table = string(Table)
	OriginalTable, _ := utils.ReadLengthCodedBinary(buffer)
	packet.OriginalTable = string(OriginalTable)
	Name, _ := utils.ReadLengthCodedBinary(buffer)
	packet.Name = string(Name)
	OriginalName, _ := utils.ReadLengthCodedBinary(buffer)
	packet.OriginalName = string(OriginalName)
	buffer.Next(1)
	packet.CharsetNumber = uint16(utils.ReadInteger(buffer, 2))
	packet.Length = uint32(utils.ReadInteger(buffer, 4))
	packet.Type = uint8(utils.ReadInteger(buffer, 1))
	packet.Flags = uint16(utils.ReadInteger(buffer, 2))
	packet.Decimals = uint8(utils.ReadInteger(buffer, 1))

	//typeOfPacket := reflect.TypeOf(*packet)
	//valueOfPacket := reflect.ValueOf(packet).Elem()
	//numField := typeOfPacket.NumField()
	//for i := 0; i < numField; i++ {
	//	field := typeOfPacket.Field(i)
	//	fieldKind := field.Type.Kind()
	//	switch fieldKind {
	//	case reflect.String:
	//		byteSlice, _ := utils.ReadLengthCodedBinary(buffer)
	//		valueOfPacket.Field(i).SetString(string(byteSlice))
	//		break
	//	case reflect.Uint8:
	//		num := utils.ReadInteger(buffer, 1)
	//		valueOfPacket.Field(i).SetUint(uint64(num))
	//		break
	//	case reflect.Uint16:
	//		num := utils.ReadInteger(buffer, 2)
	//		valueOfPacket.Field(i).SetUint(uint64(num))
	//		break
	//	case reflect.Uint32:
	//		num := utils.ReadInteger(buffer, 4)
	//		valueOfPacket.Field(i).SetUint(uint64(num))
	//		break
	//	}
	//}
}

//type ResponseEOF struct {
//	EOFMarker uint8
//	Warnings
//}

type ResponseRowData struct {
	text []string
}
