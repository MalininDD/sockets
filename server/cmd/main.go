package main

import (
	server "awesomeProject/pkg/server"
	"fmt"
)

func main() {
	serverHost := "127.0.0.1:8091"
	c := server.NewServer(serverHost)
	fmt.Println("UDP Server has been successfully started")
	c.ReadAndSendMsg()
}
