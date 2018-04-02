package handlers

import (
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
	log.Printf("received connection from %s", conn.RemoteAddr().String())

	conn.Write([]byte("connection recieved; reading message"))
	messageBytes := make([]byte, 1000)
	_, err := conn.Read(messageBytes)

	f, err := os.OpenFile(tempPath+"/command", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		errMsg := fmt.Sprintf("error creating file: %v", err)
		conn.Write([]byte(errMsg))
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
