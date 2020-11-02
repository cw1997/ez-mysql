package main

import (
	"fmt"

	"./client"
)

func main() {
	fmt.Println("ez-mysql")
	client.Client("127.0.0.1:3306", "root", "cw1997")
}
