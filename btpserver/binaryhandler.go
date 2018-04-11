package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/BKellogg/DistributedLoadTester/btp"
)

// BinaryHandler handles requests that contain binaries as the
// payload.
func BinaryHandler(w btp.ResponseWriter, r *btp.Request) {

	fp, numBytes, err := r.WritePayloadToFile("command", "")
	if err != nil {
		w.WriteString(fmt.Sprintf("error writing payload to file: %v", err))
		return
	}

	fmt.Printf("paylaod size: %d\n", r.PayloadSize)
	fmt.Printf("read %d bytes into a file\n", numBytes)

	// Executre the file at the "fp" path
	// and write its output back to the client.
	cmd := exec.Command(fp)
	cmd.Stdout = io.MultiWriter(w, os.Stdout)
	if err != nil {
		w.WriteString(fmt.Sprintf("error getting command stdout pipe: %v", err))
		return
	}
	err = cmd.Start()
	if err != nil {
		w.WriteString(fmt.Sprintf("error starting command: %v\n", err))
		return
	}
	fmt.Printf("==== Begin program output ====\n\n")
	err = cmd.Wait()
	if err != nil {
		w.WriteString(fmt.Sprintf("error waiting for command to finish: %v\n", err))
		return
	}
	fmt.Printf("\n\n==== End program output ====\n")

	// one occurred
	w.WriteString("success!")
}
