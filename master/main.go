package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

const tempPath = "/Users/Brendan/Documents/go/src/github.com/BKellogg/DistributedLoadTester/apps/helloworld/helloworld"

func main() {
	filePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(filePath)
	// Dial the agent process and get the connection
	// so we can send information to it.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("error dialing connection: %v", err)
	}
	defer conn.Close()

	// Ppen the file that we are going to be
	// sending the connection that we just opened.
	f, err := os.Open(tempPath)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	// Get the size of the file in bytes and then write those
	// to the connection so the agent process knows how many
	// bytes to read into the file on its end.
	stat, _ := f.Stat()
	err = binary.Write(conn, binary.LittleEndian, stat.Size())
	if err != nil {
		log.Fatalf("error writing filesize %s", err)
	}

	// Copy the bytes of the file into the connection
	// since the file as a source has a definite end
	// this will not block indefitely
	numBytes, err := io.Copy(conn, f)
	if err != nil {
		log.Fatalf("error reading all file bytes: %v", err)
	}

	// numBytes, err := io.Copy(conn, f)
	// if err != nil {
	// 	log.Fatalf("error copying bytes into connection: %v", err)
	// }
	log.Printf("copied %d bytes into the connection", numBytes)
	f.Close()

	// Read from the conneciton until the connection closes.
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
