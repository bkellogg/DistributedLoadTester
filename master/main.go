package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const tempPath = "/Users/Brendan/Documents/go/src/github.com/BKellogg/DistributedLoadTester/apps/helloworld/helloworld"

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("error dialing connection: %v", err)
	}

	f, err := readFile(tempPath)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	numBytes, err := io.Copy(conn, f)
	if err != nil {
		log.Fatalf("error copying bytes into connection: %v", err)
	}
	log.Printf("copied %d bytes into the connection", numBytes)
	f.Close()

	if err = printFromConnection(conn); err != nil {
		log.Fatalf("error printing from connection: %v", err)
	}
}

func readFile(path string) (*os.File, error) {
	return os.Open(path)
}

// printFromConnection reads and prints all messages from the connection
// until the process is terminated, an error occurs, or the connection is
// closed.
func printFromConnection(conn net.Conn) error {
	fmt.Printf("==== Start of connection read ====\n\n")
	for {
		responseBytes := make([]byte, 1000)
		num, err := conn.Read(responseBytes)
		if err == io.EOF {
			fmt.Printf("==== End of connection read ====")
			conn.Close()
			return nil
		}
		if err != nil {
			conn.Close()
			return err
		}
		fmt.Printf("read %d bytes\n", num)
		message := string(responseBytes)
		fmt.Printf("%s\n\n", message)
	}
}
