package handlers

import (
	"log"
	"net"
)

// PingHandler handles simple request/response
// connections.
func PingHandler(conn net.Conn) error {
	defer conn.Close()
	log.Printf("received connection from %s", conn.RemoteAddr().String())
	messageBytes := make([]byte, 1000)
	_, err := conn.Read(messageBytes)
	message := string(messageBytes)
	log.Printf("message: %v", message)
	_, err = conn.Write([]byte("pong"))
	if err != nil {
		return err
	}
	return nil
}
