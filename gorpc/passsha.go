package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash) // Convert the byte array to a hexadecimal string
}

func CheckPasswordHash(password, hash string) bool {
	return hash == HashPassword(password)
}

func main() {
	// Record start time for hashing
	start := time.Now()

	password := "secret"
	hash := HashPassword(password) // Hash the password

	// Calculate time taken for hashing
	elapsed := time.Since(start)

	// Output the results
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	fmt.Println("Match:   ", CheckPasswordHash(password, hash)) // Verifying the password
	fmt.Println("Time taken to hash password:", elapsed)
}
