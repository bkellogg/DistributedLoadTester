package server

import (
	"io"
)

// Reader returns an io.Reader that
// returns io.EOF when it reads a
// specific number of bytes, regardless
// of if there are more bytes that could be
// read.
type Reader struct {
	size   int
	reader io.Reader
}

func newReader(size int, r io.Reader) Reader {
	return Reader{size: size, reader: r}
}

// Read reads from the current reader into the destination
// byte slice until either the destination is full, or
// the number of bytes specified when the reader was
// created has been hit.
func (r Reader) Read(dst []byte) (int, error) {
	temp := make([]byte, min(r.size, len(dst)))

	numBytes, err := r.reader.Read(temp)
	if err != nil {
		return -1, err
	}
	copy(dst, temp)
	return numBytes, io.EOF
}

// min returns the lost value of the given
// integers. Needed since the math package
// doesn't implement this for ints, only
// floats.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
