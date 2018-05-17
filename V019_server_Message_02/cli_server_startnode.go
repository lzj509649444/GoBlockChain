package main

import "fmt"

func (cli *CLI) startNode() {
	nodeID := getNodeID()
	fmt.Printf("Starting node %d\n", nodeID)
	StartServer(nodeID)
}
