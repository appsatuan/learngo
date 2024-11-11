package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) string {
	// Use a fixed salt (you should generate a random salt for each password in production)
	salt := []byte("random_salt") // This should be a random salt per user in production

	// Argon2 parameters: time cost = 1, memory cost = 64KB, parallelism = 4, and hash length = 32 bytes
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return fmt.Sprintf("%x", hash)
}

func main() {
	// Measure time for hashing
	start := time.Now()

	password := "secret"
	hash := HashPassword(password)

	// Calculate elapsed time for hashing
	elapsed := time.Since(start)

	// Output results
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	fmt.Println("Time taken to hash password:", elapsed)
}
