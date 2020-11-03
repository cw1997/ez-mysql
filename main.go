package main

import (
	"fmt"

	//"./client"
	"./server"
)

func main() {
	fmt.Println("ez-mysql")
	//client.Client("127.0.0.1:3306", "root", "cw1997")
	server.Server("127.0.0.1:63306")
}
