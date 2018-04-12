package btp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

// Request represents the structure of a
// request sent to a btp server.
type Request struct {
	PayloadSize int64
	Payload     io.Reader
}

// WritePayloadToFile writes the payload in the current request
// to a file on the host machine at the given directory with the given name.
// Returns the filepath of the new file, the number of bytes written to that
// file, and an error if one occured.
// If the directory is not given, the current directory will be used.
func (r *Request) WritePayloadToFile(fileName, dir string) (string, int64, error) {
	if len(fileName) == 0 {
		return "", -1, errors.New("btp: cannot write to a file with no name")
	}
	if len(dir) == 0 {
		dir = currentDir()
	}

	f, err := os.OpenFile(dir+"/"+fileName, os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		return "", -1, fmt.Errorf("btp: error creating file: %v", err)
	}
	defer f.Close()

	// Copy the bytes from the request payload into
	// the file that was just created.
	numBytes, err := io.Copy(f, r.Payload)
	if err != nil {
		return "", 0, fmt.Errorf("btp: error copying bytes from payload to file: %v", err)
	}

	return dir + "/" + fileName, int64(numBytes), err
}

// requestFromConn returns a pointer to a Request that
// is associated with the given net.Conn.
func requestFromConn(conn net.Conn) (*Request, error) {
	payloadSize, err := nextInt64FromConn(conn)
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
	io.Writer
}

// WriteString writes the string to the given response
// writer. Returns the number bytes written or an error
// if one occurred.
func (rw ResponseWriter) WriteString(p string) (int, error) {
	return rw.Write([]byte(p))
}

// resWriterFromConn returns a ResponseWriter that
// is associated with the given net.Conn.
func resWriterFromConn(conn net.Conn) ResponseWriter {
	return ResponseWriter{Writer: conn}
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

// nextInt64FromConn reads the first 8 bytes of the connection
// into an int64 and returns it. Assumes that the first 8 bytes represent a
// valid int64. Returns an error if one occurred.
func nextInt64FromConn(conn net.Conn) (int64, error) {
	var size int64
	err := binary.Read(conn, binary.LittleEndian, &size)
	return size, err
}
