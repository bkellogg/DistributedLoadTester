package handlers

import (
	"bytes"
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
	defer conn.Close()
	log.Println("========================")
	log.Printf("received connection from %s", conn.RemoteAddr().String())

	status := fmt.Sprintf("connection recieved. creating file...\n")
	writeStatus(status, conn)

	// create the file that the executable will be written into
	f, err := os.OpenFile(tempPath+"/command", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		errMsg := fmt.Sprintf("error creating file: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}
	defer f.Close()

	status = "reading size of forthcoming file..."
	writeStatus(status, conn)

	// get the size of the executable from the first 8 bytes
	fileSizeBytes := make([]byte, 8)
	connReader := conn.(io.ReadCloser)
	_, err = connReader.Read(fileSizeBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error reading filesize: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}

	// convert the first 8 bytes into an int64
	fileSize := int64(binary.LittleEndian.Uint64(fileSizeBytes))
	status = fmt.Sprintf("%d bytes\n", fileSize)
	writeStatus(status, conn)

	status = fmt.Sprintf("reading request bytes into memory...\n")
	writeStatus(status, conn)

	// make a []byte of the size of the executable and read
	// from the connection into it
	fileBytes := make([]byte, fileSize)
	numBytes, err := conn.Read(fileBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error reading file bytes: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}
	writeStatus(fmt.Sprintf("read %d bytes into memory", numBytes), conn)

	// read the bytes from the in memory byte slice into
	// the file.
	fileBytesReader := bytes.NewReader(fileBytes)
	io.Copy(f, fileBytesReader)
	status = fmt.Sprintf("wrote %d bytes to file", numBytes)
	writeStatus(status, conn)

	_, err = conn.Write([]byte("success!"))
	if err != nil {
		return err
	}
	return nil
}

func writeStatus(status string, conn net.Conn) {
	conn.Write([]byte(status))
	log.Println(status)
}
