package client

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"../protocol"
	"../protocol/client"
	"../protocol/server"
	"../utils"
)

func Client(address string, username string, password string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("net.Dial err", err)
	}
	handleResponse(conn, username, password)
}

func handleResponse(conn net.Conn, username string, password string) {
	defer conn.Close()
	//conn.Write([]byte("welcome to my mysql"))
	//packet := protocol.ReadPacket(conn)

	header, payload, _ := protocol.ReadMySQLMessage(conn)
	fmt.Println("header", header)

	greeting := new(server.Greeting)
	greeting.Resolve(payload)
	fmt.Printf("%+v \n", greeting)

	login := new(client.Login)
	log.Println(greeting.Salt, greeting.ExtendedSalt)
	scramble := append(greeting.Salt, greeting.ExtendedSalt...)
	scramblePassword := client.MysqlNativePassword(scramble, password)
	login.ClientCapabilities = 0xa685
	login.ExtendedClientCapabilities = 0x007f
	login.MAXPacket = 1073741824
	login.Charset = 33
	login.Username = username
	login.PasswordLength = uint8(len(scramblePassword))
	login.Password = scramblePassword
	login.ClientAuthPlugin = greeting.AuthenticationPlugin
	login.ConnectionAttributes = []byte{
		0x03, 0x5f, 0x6f, 0x73, 0x05, 0x57, 0x69,
		0x6e, 0x36, 0x34, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
		0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
		0x08, 0x6c, 0x69, 0x62, 0x6d, 0x79, 0x73, 0x71,
		0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x32,
		0x35, 0x36, 0x37, 0x32, 0x07, 0x5f, 0x74, 0x68,
		0x72, 0x65, 0x61, 0x64, 0x05, 0x31, 0x36, 0x37,
		0x34, 0x30, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74,
		0x66, 0x6f, 0x72, 0x6d, 0x05, 0x41, 0x4d, 0x44,
		0x36, 0x34, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65,
		0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
		0x6f, 0x6e, 0x07, 0x31, 0x30, 0x2e, 0x31, 0x2e,
		0x32, 0x34,
	}

	fmt.Printf("login: %+v %+v \n", login, login.Build())

	sequence := header.SequenceId + 1
	protocol.WriteMySQLMessage(conn, login.Build(), sequence)

	for {
		header, payload, err := protocol.ReadMySQLMessage(conn)
		if err != nil {
			break
		}
		fmt.Printf("header: %+v , payload: %+v \n", header, payload)
		ResolveResponse(payload)
	}
}

func ResolveResponse(byteSlice []byte) {
	buffer := bytes.NewBuffer(byteSlice)
	responseCode := utils.ReadInteger(buffer, 1)
	if responseCode == server.ResponseCodeOK {
		responseOK := new(server.ResponseOK)
		responseOK.Resolve(byteSlice)
		fmt.Printf("responseOK: %+v \n", responseOK)
	} else if responseCode == server.ResponseCodeError {
		responseError := new(server.ResponseError)
		responseError.Resolve(byteSlice)
		fmt.Printf("responseError: %+v \n", responseError)
	} else {
		fmt.Printf("responseCode: %d , %+v \n", responseCode, byteSlice)
	}
}
