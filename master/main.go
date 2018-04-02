package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("error dialing connection: %v", err)
	}
	_, err = conn.Write([]byte("hello world!"))
	if err != nil {
		log.Fatalf("error writing bytes: %v", err)
	}
	responseBytes := make([]byte, 4)
	_, err = conn.Read(responseBytes)
	if err != nil {
		log.Fatalf("error reading response bytes: %v", err)
	}
	log.Printf("response: %v", string(responseBytes))
	conn.Close()
}
