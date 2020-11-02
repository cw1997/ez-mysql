package server

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

const (
	ResponseCodeOK           = 0x00
	ResponseCodeError        = 0xFF
	ResponseCodeResultStart  = 0x01
	ResponseCodeResultEnd    = 0xFA
	ResponseCodeFieldStart   = 0x01
	ResponseCodeFieldEnd     = 0xFA
	ResponseCodeRowDataStart = 0x01
	ResponseCodeRowDataEnd   = 0xFA
	ResponseCodeEOF          = 0xFE
)
