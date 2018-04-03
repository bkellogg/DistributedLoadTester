package handlers

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const tempPath = "/Users/Brendan/Documents/go/src/github.com/BKellogg/DistributedLoadTester/agent"

// ExecutableHandler handles executables
func ExecutableHandler(conn net.Conn) error {
	// Make sure the connection is not nil
	// since this will cause many things to
	// panic if this is the case.
	if conn == nil {
		return errors.New("connection must not be nil")
	}

	// close the connection when we're done with it
	defer conn.Close()

	// write some status messages to standard out and to the connection
	// so we can more easily track what is happening.
	log.Printf("received connection from %s\n", conn.RemoteAddr().String())
	write(fmt.Sprintf("connection recieved. saving request bytes as a file...\n"), conn)

	// Open or create the file that the bytes will be written to.
	// Assign the executable permission to the file to the current user.
	// This is done with the "0744" argument.
	f, err := os.OpenFile(tempPath+"/command", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		errMsg := fmt.Sprintf("error creating file: %v\n", err)
		return write(errMsg, conn)
	}
	defer f.Close()

	// Get the filesize from the connection.
	// This is done by reading the first 8 bytes
	// into an int64.
	fileSize, err := int64FromConn(conn)
	if err != nil {
		errMsg := fmt.Sprintf("error gettig int64 from conn: %s\n", err)
		return write(errMsg, conn)
	}

	// Since io.Copy has no way of knowing when to stop reading
	// we make a temp byte slice that of exactly how many bytes we need
	// so io.Copy will stop when this slice is full.
	tempBytes := make([]byte, fileSize)

	// Read the contents of the conn until the tempBytes
	// slice is full. Since the tempBytes slice was created
	// with the exact size as specifiied in the first bytes
	// of the connection it will contain exactly enough room
	// for the file.
	_, err = io.ReadFull(conn, tempBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error writing to temp storage: %v\n", err)
		return write(errMsg, conn)
	}

	// Write the contents of the tempBytes slice
	// into the file. Size this slice contains the
	// bytes of the file the client sent to us, the
	// file on the local machine will be the file that
	// the client sent to us.
	_, err = f.Write(tempBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error writing to temp storage: %v\n", err)
		return write(errMsg, conn)
	}

	// write a success message and return the error if
	// one occurred
	return write("success!", conn)
}

// write writes the given message to standard out, the connection
// given and returns the message as an error
func write(msg string, conn net.Conn) error {
	log.Print(msg)
	conn.Write([]byte(msg))
	return errors.New(msg)
}

// int64FromConn reads the first 8 bytes of the connection
// into an int64 and returns it. Assumes that the first 8 bytes represent a
// valid int64. Returns an error if one occurred.
func int64FromConn(conn net.Conn) (int64, error) {
	var size int64
	err := binary.Read(conn, binary.LittleEndian, &size)
	return size, err
}
