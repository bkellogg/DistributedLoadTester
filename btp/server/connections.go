package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// Request represents the structure of a
// request sent to a btp server.
type Request struct {
	PayloadSize int64
	Payload     io.Reader
}

// requestFromConn returns a pointer to a Request that
// is associated with the given net.Conn.
func requestFromConn(conn net.Conn) (*Request, error) {
	payloadSize, err := int64FromConn(conn)
	if err != nil {
		return nil, err
	}
	return &Request{
		PayloadSize: payloadSize,
		Payload:     newReader(payloadSize, conn),
	}, nil
}

// ResponseWriter represents the structure and
// functionality that a btp server has to respond
// to a client.
type ResponseWriter struct {
	client io.Writer
}

// Write writes the given []byte to the client.
// Returns the number of bytes written and an
// error if one occurred.
func (rw ResponseWriter) Write(p []byte) (int, error) {
	return rw.client.Write(p)
}

// resWriterFromConn returns a ResponseWriter that
// is associated with the given net.Conn.
func resWriterFromConn(conn net.Conn) ResponseWriter {
	return ResponseWriter{client: conn}
}

// fullCycleFromConn returns a ResponseWriter and a Request that
// are associated with the given conn. Returns an error if one
// occurred.
func fullCycleFromConn(conn net.Conn) (ResponseWriter, *Request, error) {
	req, err := requestFromConn(conn)
	if err != nil {
		return ResponseWriter{}, nil, fmt.Errorf("error creating request with the given conn: %v", err)
	}
	return resWriterFromConn(conn), req, nil
}

// int64FromConn reads the first 8 bytes of the connection
// into an int64 and returns it. Assumes that the first 8 bytes represent a
// valid int64. Returns an error if one occurred.
func int64FromConn(conn net.Conn) (int64, error) {
	var size int64
	err := binary.Read(conn, binary.LittleEndian, &size)
	return size, err
}
