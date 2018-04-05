package server

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
		size   int
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
