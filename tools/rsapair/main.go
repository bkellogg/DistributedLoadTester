package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BKellogg/DistributedLoadTester/shared/dir"
)

// KeyBits defines how many bits should be in the generated
// rsa public/private key pair.
const KeyBits = 2048

func main() {

	fmt.Printf("creating a %d bit rsa private/public key pair...", KeyBits)

	dir := dir.CurrentDir()
	rng := rand.Reader

	privkey, err := rsa.GenerateKey(rng, KeyBits)
	if err != nil {
		log.Fatalf("error generating rsa key pair: %v", err)
	}
	pubkey := &privkey.PublicKey

	if err = writePrivKey(privkey, dir); err != nil {
		log.Fatalf("error writing privkey to disk: %v", err)
	}
	if err = writePubKey(pubkey, dir); err != nil {
		log.Fatalf("error writing pubkey to disk %v", err)
	}
	fmt.Printf("done\n")
	fmt.Printf("private key has been written to %s\n", dir+"/privkey")
	fmt.Printf("public key has been written to %s", dir+"/pubkey")
}

// writePubKey writes the public key to a file name "pubkey"
// in the given path. Returns an error if one occurred.
func writePubKey(pubkey *rsa.PublicKey, path string) error {
	keyBytes := x509.MarshalPKCS1PublicKey(pubkey)
	f, err := os.OpenFile(path+"/pubkey", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		return errors.New("error creating pubkey file: " + err.Error())
	}
	defer f.Close()
	bytesWritten, err := f.Write(keyBytes)
	if err != nil {
		return errors.New("error writing privkey bytes to file: " + err.Error())
	}
	if bytesWritten != len(keyBytes) {
		return errors.New("number of bytes written to file does not match length of key")
	}
	return nil
}

// writePrivKey writes the private key to a file name "privkey"
// in the given path. Returns an error if one occurred.
func writePrivKey(privkey *rsa.PrivateKey, path string) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	f, err := os.OpenFile(path+"/privkey", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		return errors.New("error creating privkey file: " + err.Error())
	}
	defer f.Close()
	bytesWritten, err := f.Write(keyBytes)
	if err != nil {
		return errors.New("error writing privkey bytes to file: " + err.Error())
	}
	if bytesWritten != len(keyBytes) {
		return errors.New("number of bytes written to file does not match length of key")
	}
	return nil
}
