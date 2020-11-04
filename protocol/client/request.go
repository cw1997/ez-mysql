package client

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
	Command   uint8
	Statement string
}

func (packet *Request) Build() []byte {
	buffer := bytes.NewBuffer([]byte{})
	utils.WriteInteger(buffer, 1, uint64(packet.Command))
	buffer.WriteString(packet.Statement)
	return buffer.Bytes()
}

func (packet *Request) Resolve(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	packet.Command = uint8(utils.ReadInteger(buffer, 1))
	packet.Statement = string(buffer.Next(len(byteSlice)))
}

const (
	COM_SLEEP               = 0x00 //
	COM_QUIT                = 0x01 //	mysql_close
	COM_INIT_DB             = 0x02 //	mysql_select_db
	COM_QUERY               = 0x03 //	mysql_real_query
	COM_FIELD_LIST          = 0x04 //	mysql_list_fields
	COM_CREATE_DB           = 0x05 //	mysql_create_db
	COM_DROP_DB             = 0x06 //	mysql_drop_db
	COM_REFRESH             = 0x07 //	mysql_refresh
	COM_SHUTDOWN            = 0x08 //	mysql_shutdown
	COM_STATISTICS          = 0x09 //	mysql_stat
	COM_PROCESS_INFO        = 0x0A //	mysql_list_processes
	COM_CONNECT             = 0x0B //
	COM_PROCESS_KILL        = 0x0C //	mysql_kill
	COM_DEBUG               = 0x0D //	mysql_dump_debug_info
	COM_PING                = 0x0E //	mysql_ping
	COM_TIME                = 0x0F //
	COM_DELAYED_INSERT      = 0x10 //
	COM_CHANGE_USER         = 0x11 //	mysql_change_user
	COM_BINLOG_DUMP         = 0x12 //
	COM_TABLE_DUMP          = 0x13 //
	COM_CONNECT_OUT         = 0x14 //
	COM_REGISTER_SLAVE      = 0x15 //
	COM_STMT_PREPARE        = 0x16 //	mysql_stmt_prepare
	COM_STMT_EXECUTE        = 0x17 //	mysql_stmt_execute
	COM_STMT_SEND_LONG_DATA = 0x18 //	mysql_stmt_send_long_data
	COM_STMT_CLOSE          = 0x19 //	mysql_stmt_close
	COM_STMT_RESET          = 0x1A //	mysql_stmt_reset
	COM_SET_OPTION          = 0x1B //	mysql_set_server_option
	COM_STMT_FETCH          = 0x1C //	mysql_stmt_fetch
)
