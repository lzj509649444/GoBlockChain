package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

const protocol = "tcp"
const commandLength = 12

var nodeAddress string

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func handleConnection(conn net.Conn) {
	requestBytes, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Panic(err)
	}
	command := bytesToCommand(requestBytes[:commandLength])

	switch command {
	case "version":
		fmt.Printf("Received %s command\n", command)
	default:
		fmt.Println("Unknown command received!")
	}

	// Send a response back to person contacting us.
	conn.Write([]byte(fmt.Sprintf("Received %s !\n", requestBytes)))

	conn.Close()

}

// StartServer starts a node
func StartServer(nodeID int) {
	nodeAddress = fmt.Sprintf("localhost:%d", nodeID)
	listen, err := net.Listen(protocol, nodeAddress)

	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn)
	}

}
