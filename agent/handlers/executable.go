package handlers

import (
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

	conn.Write([]byte(fmt.Sprintf("connection recieved. creating file...\n")))

	f, err := os.OpenFile(tempPath+"/command", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		errMsg := fmt.Sprintf("error creating file: %v", err)
		conn.Write([]byte(errMsg))
		return errors.New(errMsg)
	}
	defer f.Close()

	conn.Write([]byte(fmt.Sprintf("copying request bytes into file...\n")))

	fmt.Println("copying bytes into file")
	numBytes, err := io.Copy(f, conn)
	if err != nil {
		fmt.Println("error while copying: " + err.Error())
		errMsg := fmt.Sprintf("error writing to file: %v", err)
		conn.Write([]byte(errMsg))
		return errors.New(errMsg)
	}
	fmt.Println("dont copying bytes into file")
	conn.Write([]byte(fmt.Sprintf("wrote %d bytes to file", numBytes)))

	// cmd := exec.Command(tempPath + "/command")
	// err = cmd.Run()
	// if err != nil {
	// 	errMsg := fmt.Sprintf("error executing command: %v", err)
	// 	conn.Write([]byte(errMsg))
	// 	return errors.New(errMsg)
	// }

	_, err = conn.Write([]byte("success!"))
	if err != nil {
		return err
	}
	return nil
}
