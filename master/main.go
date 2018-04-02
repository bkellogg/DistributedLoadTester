package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("error dialing connection: %v", err)
	}
	_, err = conn.Write([]byte("hello world!"))
	if err != nil {
		log.Fatalf("error writing bytes: %v", err)
	}
	if err = printFromConnection(conn); err != nil {
		log.Fatalf("error printing from connection: %v", err)
	}
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
