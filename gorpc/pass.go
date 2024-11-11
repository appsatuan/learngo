package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	// Record start time for hashing
	start := time.Now()

	password := "secret"
	hash, _ := HashPassword(password) // ignore error for simplicity

	// Calculate time taken for hashing
	elapsed := time.Since(start)

	// Output the results
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	fmt.Println("Match:   ", CheckPasswordHash(password, hash))
	//elapsed2 := time.Since(start)
	fmt.Println("Time taken to hash password:", elapsed)
}
