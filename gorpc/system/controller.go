package system

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	DB *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{DB: db}
}

func (c *Controller) GetLoaders(ctx *gin.Context) {
	rows, err := c.DB.Query("SELECT id, value, timestamp FROM loader")
	if err != nil {
		log.Fatal(err)
		ctx.JSON(500, gin.H{"error": "Database query failed"})
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

	ctx.JSON(200, gin.H{"loaders": loaders})
}
