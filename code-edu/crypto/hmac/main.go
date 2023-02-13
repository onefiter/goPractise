package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {

	msg := "This is the message to hash!"

	// hmac
	hmacData := hmacHash(msg, "This is the key!")
	fmt.Printf("HMAC: %x\n", hmacData)

}

func hmacHash(msg string, key string) (hashData []byte) {
	k := []byte(key)
	mac := hmac.New(sha1.New, k)
	io.WriteString(mac, msg)
	hashData = mac.Sum(nil)
	return
}
