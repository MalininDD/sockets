package client

import (
	"encoding/json"
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

type Client struct {
	Name string
	Host string
	Conn *net.UDPConn
}

type Msg struct {
	Name  string `json:"name"`
	Voice []byte `json:"voice"`
}


func NewClient(Name string, Host string) Client {
	s, err := net.ResolveUDPAddr("udp4", Host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Port, s.IP, s.Zone)

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		log.Fatal(err)
	}
	return Client{Name: Name, Host: Host, Conn: c}
}

func (c *Client) SendVoiceMsg(filePath string) {



	fileByte, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	msg := Msg{Name: c.Name, Voice: fileByte}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.Conn.Write(jsonMsg)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func (c *Client) ReadMsgAndPlay() {
	buffer := make([]byte, 60000)
	n, _, err := c.Conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	}

	var input Msg
	err = json.Unmarshal(buffer[:n], &input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully received msg from ", input.Name)
	err = ioutil.WriteFile("voice.mp3", input.Voice, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("voice.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
	speaker.Play(streamer)
	time.Sleep(1 * time.Second)
	err = c.Conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success")
}
