package handlers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

const tempPath = "/Users/Brendan/Documents/go/src/github.com/BKellogg/DistributedLoadTester/agent"

// ExecutablenHandler handles executables
func ExecutableHandler(conn net.Conn) error {
	defer conn.Close()
	log.Println("========================")
	log.Printf("received connection from %s", conn.RemoteAddr().String())

<<<<<<< HEAD
	status := fmt.Sprintf("connection recieved. creating file...\n")
	writeStatus(status, conn)
=======
	conn.Write([]byte("connection recieved; reading message"))
	messageBytes := make([]byte, 1000)
	_, err := conn.Read(messageBytes)
>>>>>>> parent of daf31cf... can send executables over tcp connection

	// create the file that the executable will be written into
	f, err := os.OpenFile(tempPath+"/command", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		errMsg := fmt.Sprintf("error creating file: %v", err)
		writeStatus(errMsg, conn)
		return errors.New(errMsg)
	}
	defer f.Close()

	numBytes, err := f.Write(messageBytes)
	if err != nil {
		errMsg := fmt.Sprintf("error writing to file: %v", err)
		conn.Write([]byte(errMsg))
		return errors.New(errMsg)
	}
	conn.Write([]byte(fmt.Sprintf("wrote %d bytes to file", numBytes)))

	cmd := exec.Command(tempPath + "/command")
	err = cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("error executing command: %v", err)
		conn.Write([]byte(errMsg))
		return errors.New(errMsg)
	}

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
