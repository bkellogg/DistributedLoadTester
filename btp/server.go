package btp

import (
	"fmt"
	"log"
	"net"
)

// Listen starts a btp server listening on the given
// address and handles all requests with the given
// handler. Returns an error if one occurred.
//
// Only errors encountered during startup will be returned.
// Errors encountered during while processing a specific
// connection during any point in it's lifecycle will not
// be returned here.
//
// This function is a blocking function and will never exit
// once properly started.
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
		go serveRequest(conn, handler)
	}
}

// serveRequest serves the given connection with the
// given handler. Closes the connection when the serving
// is complete.
func serveRequest(conn net.Conn, handler Handler) {
	w, r, err := fullCycleFromConn(conn)
	if err != nil {
		log.Printf("error getting conn lifecycle: %v", err)
		conn.Close()
	}
	handler(w, r)
	conn.Close()
}
