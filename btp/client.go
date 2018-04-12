package btp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

// RequestBuilder represents a builder
// for BTP requests
type RequestBuilder struct {
	hasPayload  bool
	beenSent    bool
	address     string
	file        string
	payloadSize int64
	response    io.ReadCloser
}

// NewRequestBuilder returns an empty RequestBuilder
// pointed at the given address.
func NewRequestBuilder(address string) *RequestBuilder {
	return &RequestBuilder{hasPayload: false, address: address}
}

// SetFile sets the given file path as the file
// to send when the request is sent.
func (rb *RequestBuilder) SetFile(filepath string) {
	rb.hasPayload = true
	rb.file = filepath
}

// Send sends the RequestBuilder's request.
// Returns an error if one occurred.
func (rb *RequestBuilder) Send() error {
	if !rb.hasPayload {
		return fmt.Errorf("btp: cannot send a request with no body")
	}
	if rb.beenSent {
		return fmt.Errorf("btp: cannot sent a request that has already been sent")
	}

	// open the file of this request builder
	// report any errors that occur.
	f, err := os.Open(rb.file)
	if err != nil {
		return fmt.Errorf("btp: error opening file: %v", err)
	}

	// obtain the file statistic so we can get the size of the file
	// report any errors that occur.
	fstat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("btp: error obtaining file statistics: %v", err)
	}

	// open the connection to this request builder's
	// address. Report any errors that occur.
	conn, err := net.Dial("tcp", rb.address)
	if err != nil {
		return fmt.Errorf("btp: error dialing address: %v", rb.address)
	}

	// write the size of the file to the connection
	// report any errors that occur.
	if err = binary.Write(conn, binary.LittleEndian, fstat.Size()); err != nil {
		conn.Close()
		return fmt.Errorf("btp: error writing file size: %v", err)
	}

	// copy the contents of the file into the connection
	// report any errors that occur.
	if _, err = io.Copy(conn, f); err != nil {
		conn.Close()
		return fmt.Errorf("btp: error copying file into connection: %v", err)
	}

	rb.response = conn
	rb.beenSent = true

	return nil
}

// Response gets the response readCloser from the connection.
// Returns an error if one occurred.
func (rb *RequestBuilder) Response() (io.ReadCloser, error) {
	if !rb.beenSent {
		return nil, fmt.Errorf("btp: cannot get the response of a request that has not been sent")
	}
	if rb.response == nil {
		return nil, fmt.Errorf("btp: response is nil")
	}
	return rb.response, nil
}
