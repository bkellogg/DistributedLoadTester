package main

import (
	"log"

	"github.com/BKellogg/DistributedLoadTester/btp"
)

func main() {
	addr := "localhost:8080"
	log.Printf("btp listening on %s...\n", addr)
	log.Fatal(btp.Listen(addr, BinaryHandler))
}
