package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/BKellogg/DistributedLoadTester/btp"
)

const poorSucker = "localhost:8080"

func main() {

	if len(os.Args) < 2 {
		log.Fatal(`usage:
			btpclient <app_path>`)
	}
	appPath := os.Args[1]
	fullPath, err := filepath.Abs(appPath)
	if err != nil {
		log.Fatalf("error getting absolute path: %v", err)
	}

	reqBuilder := btp.NewRequestBuilder(poorSucker)
	reqBuilder.SetFile(fullPath)
	if err := reqBuilder.Send(); err != nil {
		log.Fatalf("error sending request: %v", err)
	}

	response, err := reqBuilder.Response()
	if err != nil {
		log.Fatalf("error getting response: %v", err)
	}
	defer response.Close()

	// Read from the conneciton until the connection closes.
	if err = copyFromRCIntoStdOut(response); err != nil {
		log.Fatalf("error printing from readcloser: %v", err)
	}
}

// copyRCConnIntoStdOut reads and prints all messages from the readcloser
// until the process is terminated, an error occurs, or the connection is
// closed.
func copyFromRCIntoStdOut(rc io.ReadCloser) error {
	fmt.Printf("==== Start of connection read ====\n\n")
	numBytes, err := io.Copy(os.Stdout, rc)
	fmt.Printf("\n\n==== End of connection read ====\n")
	fmt.Printf("total bytes read: %d", numBytes)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error copying into connection into standard out: %v", err)
	}
	return nil
}
