package main

import (
	"log"

	"gorpc/system"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := system.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := system.NewRouter(db)
	router.Static("/frontend", "./frontend")
	// Start the server
	router.Run(":1234")
}
