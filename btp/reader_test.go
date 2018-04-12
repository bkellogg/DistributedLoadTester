package btp

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestNewReader(t *testing.T) {
	reader1 := bytes.NewReader([]byte("some string"))
	reader2 := ioutil.NopCloser(reader1)

	cases := []struct {
		size   int64
		reader io.Reader
	}{
		{
			1,
			reader1,
		},
		{
			18329320,
			reader2,
		},
	}
	for _, c := range cases {
		r := newReader(c.size, c.reader)
		if r.size != c.size {
			t.Fatalf("error in NewReader: sizes %d and %d did not match", r.size, c.size)
		}
	}
}

func TestRead(t *testing.T) {
	cases := []struct {
		reader        *Reader
		expectedBytes []byte
	}{
		{
			newReader(10, bytes.NewReader([]byte("hi there!!"))),
			[]byte("hi there!!"),
		},
		{
			newReader(100, bytes.NewReader([]byte("hi there!!"))),
			[]byte("hi there!!"),
		},
		{
			newReader(10, bytes.NewReader([]byte("hi there!! how are you doing today?"))),
			[]byte("hi there!!"),
		},
		{
			newReader(1000, bytes.NewReader([]byte("hi there!! how are you doing today?"))),
			[]byte("hi there!! how are you doing today?"),
		},
		{
			newReader(0, bytes.NewReader([]byte("hi there!! how are you doing today?"))),
			[]byte(""),
		},
		{
			newReader(4, bytes.NewReader([]byte("pink fluffy ponies make me happy"))),
			[]byte("pink"),
		},
	}
	for _, c := range cases {
		dst := make([]byte, len(c.expectedBytes))
		numRead, err := c.reader.Read(dst)
		if err != nil {
			if err == io.EOF {
				continue
			}
			t.Fatalf("error reading from Reader: %v", err)
		}
		if numRead != len(c.expectedBytes) {
			t.Fatalf("%d bytes read did not match expected number %d", numRead, len(c.expectedBytes))
		}
		stringIn, stringOut := string(c.expectedBytes), string(dst)
		if stringIn != stringOut {
			t.Fatalf("input %s did not match %s output", stringIn, stringOut)
		}

	}
}
