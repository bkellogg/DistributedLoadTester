package server

import "net"

// Handler defines the type of function that can be
// used as a BTP Handler
type Handler func(conn net.Conn)
