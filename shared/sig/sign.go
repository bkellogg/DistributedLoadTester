package sig

import (
	"crypto/rsa"
	"io"
)

// SignFromReaderToWriter signs the bytes from src io.Reader and writes them
// to the dst io.Writer using the given PrivateKey privkey. Returns an error
// if one occurred.
func SignFromReaderToWriter(src io.Reader, dst io.Writer, privkey *rsa.PrivateKey) error {
	panic("TODO")
}
