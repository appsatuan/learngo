package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
			return
		}

		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer openedFile.Close()

		start := time.Now()
		f, err := excelize.OpenReader(openedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel file"})
			return
		}

		totalLines := 0
		totalWords := 0

		for _, sheet := range f.GetSheetList() {
			rows, err := f.GetRows(sheet)
			if err != nil {
				log.Printf("Failed to read rows from sheet %s: %v", sheet, err)
				continue
			}
			for _, row := range rows {
				totalLines++
				for _, cell := range row {
					words := strings.Fields(cell)
					totalWords += len(words)
				}
			}
		}

		elapsed := time.Since(start)
		c.JSON(http.StatusOK, gin.H{
			"total_lines":     totalLines,
			"total_words":     totalWords,
			"processing_time": elapsed.Seconds(),
		})
	})

	r.Run(":1234")
}
