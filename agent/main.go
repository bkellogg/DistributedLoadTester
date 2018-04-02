package main

import (
	"log"

	"github.com/BKellogg/DistributedLoadTester/agent/handlers"
	"github.com/BKellogg/DistributedLoadTester/agent/listen"
)

const defaultListenAddr = ":8080"

func main() {
	log.Printf("listening on tcp:%s", defaultListenAddr)
	log.Fatal(listen.Listen(defaultListenAddr, handlers.ConnectionHandler))
}
