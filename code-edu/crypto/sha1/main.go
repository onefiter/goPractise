package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	msg := "This is the message to hash!"

	// sha1 
	sha1Data := sha1Hash(msg)
	fmt.Printf("SHA1:%x\n", sha1Data)
}

func sha1Hash(msg string) (hashData []byte) {
	h := sha1.New()
	io.WriteString(h, msg)
	hashData = h.Sum(nil)
	return
}