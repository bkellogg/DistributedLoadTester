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

// const poorSucker = "172.28.102.87:8080"

const poorSucker = "localhost:8080"

func main() {

	if len(os.Args) < 2 {
		log.Fatal(`usage:
			master <app_path>`)
	}
	appPath := os.Args[1]
	fullPath, err := filepath.Abs(appPath)
	if err != nil {
		log.Fatalf("error getting absolute path: %v", err)
	}
	// Dial the agent process and get the connection
	// so we can send information to it.
	conn, err := net.Dial("tcp", poorSucker)
	if err != nil {
		log.Fatalf("error dialing connection: %v", err)
	}
	defer conn.Close()

	// Ppen the file that we are going to be
	// sending the connection that we just opened.
	f, err := os.Open(fullPath)
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
	if err = copyFromConnIntoStdOut(conn); err != nil {
		log.Fatalf("error printing from connection: %v", err)
	}
}

// copyFromConnIntoStdOut reads and prints all messages from the connection
// until the process is terminated, an error occurs, or the connection is
// closed.
func copyFromConnIntoStdOut(conn net.Conn) error {
	fmt.Printf("==== Start of connection read ====\n\n")
	numBytes, err := io.Copy(os.Stdout, conn)
	fmt.Printf("\n\n==== End of connection read ====\n")
	fmt.Printf("total bytes read: %d", numBytes)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error copying into connection into standard out: %v", err)
	}
	return nil
}
