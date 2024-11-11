package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func writeSequentially(fileName string, rows int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < rows; i++ {
		record := []string{strconv.Itoa(i), "Name_" + strconv.Itoa(i)}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	rows := 1000000
	start := time.Now()

	err := writeSequentially("output_sequential.csv", rows)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Sequential write complete in %s.\n", elapsed)
}
