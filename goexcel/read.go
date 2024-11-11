package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Define a struct to hold the data from the `loader` table
type Loader struct {
	ID        int    `json:"id"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

func main() {
	// Set up the router
	r := gin.Default()

	// Set up the MySQL connection
	dsn := "web:xcpass@tcp(localhost:3306)/dbci3"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Route to handle fetching data from the `loader` table
	r.GET("/", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, value, timestamp FROM loader")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
			return
		}
		defer rows.Close()

		var loaders []Loader
		for rows.Next() {
			var loader Loader
			if err := rows.Scan(&loader.ID, &loader.Value, &loader.Timestamp); err != nil {
				log.Fatal(err)
			}
			loaders = append(loaders, loader)
		}

		// Render data in an HTML template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"loaders": loaders,
		})
	})

	// Set up the HTML templates (You can create a folder named `templates`)
	r.LoadHTMLGlob("templates/*")

	// Start the server
	r.Run(":1234")
}
