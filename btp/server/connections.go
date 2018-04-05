package server

import (
	"encoding/binary"
	"io"
	"net"
)

type Request struct {
	PayloadSize int
	Body        io.Reader
}

func requestFromConn(conn net.Conn) (*Request, error) {
	return nil, nil
}

type ResponseWriter struct {
}

// int64FromConn reads the first 8 bytes of the connection
// into an int64 and returns it. Assumes that the first 8 bytes represent a
// valid int64. Returns an error if one occurred.
func int64FromConn(conn net.Conn) (int64, error) {
	var size int64
	err := binary.Read(conn, binary.LittleEndian, &size)
	return size, err
}
