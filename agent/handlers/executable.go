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
	log.Printf("received connection from %s", conn.RemoteAddr().String())

	status := fmt.Sprintf("connection recieved. creating file...\n")
	writeStatus(status, conn)

	f, err := os.OpenFile(tempPath+"/command", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		errMsg := fmt.Sprintf("error creating file: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}
	defer f.Close()

	status = "reading size of forthcoming file..."
	writeStatus(status, conn)

	fileSizeBytes := make([]byte, 8)

	connReader := conn.(io.ReadCloser)
	_, err = connReader.Read(fileSizeBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error reading filesize: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}

	fileSize := int64(binary.LittleEndian.Uint64(fileSizeBytes))
	status = fmt.Sprintf("%d bytes\n", fileSize)
	writeStatus(status, conn)

	status = fmt.Sprintf("reading request bytes into memory...\n")
	writeStatus(status, conn)

	fileBytes := make([]byte, fileSize)
	numBytes, err := connReader.Read(fileBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error reading file bytes: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}

	writeStatus(fmt.Sprintf("read %d bytes into memory", numBytes), conn)

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
