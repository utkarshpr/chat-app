package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateStrongSecretKey(length int) string {
	// Create a random byte slice of the given length
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatal("Error generating random bytes:", err)
	}

	// Return the base64 encoded string of the random bytes
	return base64.StdEncoding.EncodeToString(secret)
}
