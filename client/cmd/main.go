package main

import (
	client "awesomeProject/pkg/client"
)



func main() {
	clientName := "Namell"
	serverHost := "127.0.0.1:8091"
	c := client.NewClient(clientName, serverHost)
	c.SendVoiceMsg("./voice.mp3")
	c.ReadMsgAndPlay()
}
