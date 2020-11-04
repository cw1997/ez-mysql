package server

import (
	"fmt"
	"log"
	"net"

	"../protocol"
	"../protocol/client"
	"../protocol/server"
)

func Server(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Error listening: ", err)
	}
	defer listener.Close()
	fmt.Println("Listening on " + address)

	threadId := uint32(0)

	for {
		threadId++
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting: ", err)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn, threadId)
	}
}

func handleRequest(conn net.Conn, threadId uint32) {
	defer conn.Close()

	greeting := new(server.Greeting)
	greeting.Protocol = 10
	greeting.Version = Version // "5.7.30-log"
	greeting.ThreadId = threadId
	salt := builderSalt()
	greeting.Salt = salt[:8]
	greeting.ServerCapabilities = 0xffff
	greeting.ServerLanguage = 8
	greeting.ServerStatus = 0x0002
	greeting.ExtendedServerStatus = 0xc1ff
	greeting.ExtendedSaltLength = 21
	greeting.ExtendedSalt = salt[8:]
	greeting.AuthenticationPlugin = "mysql_native_password"
	greetingBytes := greeting.Build()

	n, err := protocol.WriteMySQLMessage(conn, greetingBytes, 0)
	if err != nil {
		log.Print("protocol.WriteMySQLMessage(conn, greeting.Build())", n, err)
		return
	}

	header, payload, err := protocol.ReadMySQLMessage(conn)
	if err != nil {
		return
	}
	login := new(client.Login)
	login.Resolve(payload)

	responseOK := new(server.ResponseOK)
	responseOK.ResponseCode = server.ResponseCodeOK
	responseOK.AffectedRows = 0
	responseOK.LastInsertID = 0
	responseOK.ServerStatus = 0x0002
	responseOK.Warnings = 0

	sequenceId := header.SequenceId + 1
	n, err = protocol.WriteMySQLMessage(conn, responseOK.Build(), sequenceId)
	if err != nil {
		log.Print("protocol.WriteMySQLMessage(conn, responseOK.Build())", n, err)
		return
	}

	for {
		header, payload, err := protocol.ReadMySQLMessage(conn)
		if err != nil {
			break
		}
		fmt.Printf("header: %+v , payload: %+v \n", header, payload)
		fmt.Printf("header: %+v \n", header)
		request := new(client.Request)
		request.Resolve(payload)
		switch request.Command {
		case client.COM_QUERY:
			fieldNum := 1
			sequenceId = 1
			protocol.WriteMySQLMessage(conn, []byte{byte(fieldNum)}, sequenceId)

			field := new(server.ResponseField)
			field.Catalog = "def"
			field.Database = ""
			field.Table = ""
			field.OriginalTable = ""
			field.Name = "@@version_comment"
			field.OriginalName = "@@version_comment"
			field.CharsetNumber = 33
			field.Length = 84
			field.Type = 253
			field.Flags = 0x0000
			field.Decimals = 31
			sequenceId++
			protocol.WriteMySQLMessage(conn, field.Build(), sequenceId)

			responseEOFField := new(server.ResponseEOF)
			responseEOFField.Warnings = 0
			responseEOFField.ServerStatus = 0x0002
			sequenceId++
			protocol.WriteMySQLMessage(conn, responseEOFField.Build(), sequenceId)

			rowData := new(server.ResponseRowData)
			rowData.Text = append(rowData.Text, VersionComment)
			sequenceId++
			protocol.WriteMySQLMessage(conn, rowData.Build(), sequenceId)

			responseEOFRow := new(server.ResponseEOF)
			responseEOFRow.Warnings = 0
			responseEOFRow.ServerStatus = 0x0002
			sequenceId++
			protocol.WriteMySQLMessage(conn, responseEOFRow.Build(), sequenceId)
		}
	}
}

func builderSalt() []byte {
	return []byte{
		0x6c, 0x72, 0x3b, 0x26, 0x15, 0x3f, 0x4b, 0x2a,
		0x50, 0x40, 0x67, 0x18, 0x6c, 0x48, 0x2a, 0x4f,
		0x72, 0x6e, 0x2c, 0x11,
	}
}
