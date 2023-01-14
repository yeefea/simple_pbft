package main

import (
	"os"

	"github.com/yeefea/simple_pbft/pbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
