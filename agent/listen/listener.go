package listen

import (
	"fmt"
	"log"
	"net"
)

// Listen listens on the given addr and sends the connections
// to the given handler func.
func Listen(addr string, handler func(net.Conn) error) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening in %s: %v", addr, err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			// no need to stop listening even if a connection failed to open
			log.Printf("error accepting connection: %v\n", err)
		}
		go handler(conn)
	}
}
