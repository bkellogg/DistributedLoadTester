package btp

import (
	"bytes"
	"testing"
)

func TestRWWrite(t *testing.T) {
	cases := []struct {
		payload []byte
	}{
		{
			[]byte("hello there"),
		},
		{
			[]byte("i love playing with ponies"),
		},
		{
			[]byte("stable genius"),
		},
		{
			[]byte("i am like, very smart"),
		},
		{
			[]byte("*@&!)!@*(@#(dkjfhs7aksdj??sd'$%"),
		},
		{
			[]byte(""),
		},
	}
	for _, c := range cases {
		var b bytes.Buffer
		rw := ResponseWriter{client: &b}
		payloadLength := len(c.payload)

		numWritten, err := rw.Write(c.payload)
		if err != nil {
			t.Fatalf("error writing payload to ResponseWriter: %v", err)
		}
		if numWritten != payloadLength {
			t.Fatalf("number bytes written %d did not match expected number of bytes written %d",
				numWritten, payloadLength)
		}
		output := make([]byte, payloadLength)
		_, err = b.Read(output)
		if err != nil {
			t.Fatalf("error reading from buffer: %v", err)
		}
		if !bytes.Equal(c.payload, output) {
			t.Fatal("payload bytes did not match the returned output bytes")
		}
	}
}
