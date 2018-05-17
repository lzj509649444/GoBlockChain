package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

const nodeVersion = 1

var knownNodes []string

// verack The verack message is sent in reply to version. This message consists of only a message header with the command string "verack"
type verack struct {
}

// addr: Provide information on known nodes of the network
type addr struct {
	AddrList []string
}

// When a node creates an outgoing connection, it will immediately advertise its version. The remote node will respond with its version. No further communication is possible until both peers have exchanged their version
// version 用于找到一个更长的区块链。当一个节点接收到 version 消息，它会检查本节点的区块链是否比 BestHeight 的值更大。如果不是，节点就会请求并下载缺失的块。
type verzion struct {
	Version int

	AddrFrom string
}

func sendVrack(addr string) {
	payload := gobEncode(verack{})

	request := append(commandToBytes("verack"), payload...)

	sendData(addr, request)
}

func sendAddr(address string) {
	nodes := addr{knownNodes}
	nodes.AddrList = append(nodes.AddrList, nodeAddress)
	payload := gobEncode(nodes)
	request := append(commandToBytes("addr"), payload...)

	sendData(address, request)
}

func handleAddr(request []byte) {
	var buff bytes.Buffer
	var payload addr

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	knownNodes = append(knownNodes, payload.AddrList...)
	fmt.Printf("There are %d known nodes now!\n", len(knownNodes))
}

func sendVersion(addr string) {
	payload := gobEncode(verzion{nodeVersion, nodeAddress})

	request := append(commandToBytes("version"), payload...)

	sendData(addr, request)
}

func handleVersion(request []byte) {
	var buff bytes.Buffer
	var payload verzion

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	sendVrack(payload.AddrFrom)
	sendAddr(payload.AddrFrom)
	knownNodes = append(knownNodes, payload.AddrFrom)
}
