package utils

import (
	"crypto/rand"
	"log"
)

func RandomBytes(bytesLength int) []byte {
	buf := make([]byte, 128)
	// then we can call rand.Read.
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}

	return buf
}
