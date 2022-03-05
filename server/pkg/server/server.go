package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Msg struct {
	Name  string `json:"name"`
	Voice []byte `json:"voice"`
}

type Server struct {
	Host string
	Conn *net.UDPConn
}

func NewServer(Host string) Server {
	s, err := net.ResolveUDPAddr("udp4", Host)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatal(err)
	}

	return Server{Host: Host, Conn: conn}
}

func (s *Server) ReadAndSendMsg() {
	buffer := make([]byte, 60000)

	for {
		n, addr, err := s.Conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		var input Msg
		err = json.Unmarshal(buffer[:n], &input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully received msg from", input.Name)
		fmt.Println("Sending...")
		_, err = s.Conn.WriteToUDP(buffer[:n], addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Success")
		//break
	}


}