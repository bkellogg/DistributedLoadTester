package server

import (
	"io"
)

// Reader returns an io.Reader that
// returns io.EOF when it reads a
// specific number of bytes, regardless
// of if there are more bytes that could be
// read. The client of this reader will be
// unable to change the size of the reader
// so setting it to the size of the body is
// fine.
type Reader struct {
	size   int64
	reader io.Reader
}

// Read reads from the current reader into the destination
// byte slice until either the destination is full, or
// the number of bytes specified when the reader was
// created has been hit.
func (r Reader) Read(dst []byte) (int, error) {
	temp := make([]byte, min(r.size, int64(len(dst))))

	numBytes, err := r.reader.Read(temp)
	if err != nil {
		return -1, err
	}
	copy(dst, temp)
	return numBytes, io.EOF
}

// newReader creates a new Reader that wraps the given
// reader. The returned reader will return io.EOF when
// the source reader is exhausted or the number of bytes
// read is equal to the given size.
func newReader(size int64, r io.Reader) Reader {
	return Reader{size: size, reader: r}
}

// min returns the lost value of the given
// integers. Needed since the math package
// doesn't implement this for ints, only
// floats.
func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
