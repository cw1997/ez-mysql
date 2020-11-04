package server

import (
	"../../utils"
	"bytes"
)

type ResponseResultSet struct {
	ResponseField   []ResponseField
	ResponseRowData []ResponseRowData
}

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
