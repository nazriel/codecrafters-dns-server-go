package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/app/dns"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		requestMsg := dns.MessageFromBytes(buf[:size])
		if requestMsg == nil {
			fmt.Println("Failed to parse request")
			continue
		}

		questions := []dns.Question{}
		questions = append(questions, dns.Question{Name: dns.NameFromString("codecrafters.io"), Type: 1, Class: 1})

		answers := []dns.Answer{}
		answers = append(answers, dns.Answer{Name: dns.NameFromString("codecrafters.io"), Type: 1, Class: 1, TTL: 3600, Data: dns.IpFromString("1.1.1.1")})

		response := dns.Message{
			Header: dns.Header{
				PacketID: requestMsg.Header.PacketID,
				HeaderFlags: dns.HeaderFlags{
					QueryResponseIndiciator: 1,
					Opcode:                  requestMsg.Header.HeaderFlags.Opcode,
					RecursionDesired:        requestMsg.Header.HeaderFlags.RecursionDesired,
					ResponseCode: func() uint8 {
						if requestMsg.Header.HeaderFlags.Opcode == 0 {
							return 0
						} else {
							return 4
						}
					}(),
				},
				QuestionCount:     uint16(len(questions)),
				AnswerRecordCount: uint16(len(answers)),
			},
			Questions: questions,
			Answers:   answers,
		}

		_, err = udpConn.WriteToUDP(response.Bytes(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
