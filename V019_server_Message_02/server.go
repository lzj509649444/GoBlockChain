package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

const protocol = "tcp"
const commandLength = 12
const dnsNodeID = 4000

var nodeAddress string

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func commandToBytes(command string) []byte {
	var bytes [commandLength]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func sendData(addr string, data []byte) {
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}

func handleConnection(conn net.Conn, blockchain *Blockchain) {
	requestBytes, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Panic(err)
	}
	command := bytesToCommand(requestBytes[:commandLength])
	fmt.Printf("Received %s command\n", command)

	switch command {
	case "addr":
		handleAddr(requestBytes)
	case "version":
		handleVersion(requestBytes, blockchain)
	case "verack":
		//
	case "block":
		handleBlock(requestBytes, blockchain)
	case "inv":
		handleInv(requestBytes, blockchain)
	case "getblocks":
		handleGetBlocks(requestBytes, blockchain)
	case "getdata":
		handleGetData(requestBytes, blockchain)
	default:
		fmt.Println("Unknown command!")
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

	blockchain := GetBlockchain()

	// if nodeID != dnsNodeID {
	// 	sendVersion(fmt.Sprintf("localhost:%d", dnsNodeID), blockchain)
	// }

	if nodeAddress != knownNodes[0] {
		sendVersion(knownNodes[0], blockchain)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn, blockchain)
	}

}
