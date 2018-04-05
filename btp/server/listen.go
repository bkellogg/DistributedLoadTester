package server

import (
	"fmt"
	"log"
	"net"
)

// Listen begins listening for requests on the given
// addr and sends all requests to the given handler.
func Listen(addr string, handler Handler) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening in %s: %v", addr, err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error accepting connection: %v\n", err)
		}
		w, r, err := fullCycleFromConn(conn)
		if err != nil {
			log.Printf("error getting conn lifecycle: %v", err)
			conn.Close()
			continue
		}
		go handler(w, r)
	}
}
