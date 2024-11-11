package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func writeChunk(writer *csv.Writer, start int, end int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	// Collect records and write them in a single batch
	var records [][]string
	for i := start; i < end; i++ {
		records = append(records, []string{strconv.Itoa(i), "Name_" + strconv.Itoa(i)})
	}

	mu.Lock()
	if err := writer.WriteAll(records); err != nil {
		fmt.Println("Error writing records:", err)
	}
	mu.Unlock()
}

func writeConcurrently(fileName string, rows, chunkSize, maxGoroutines int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, maxGoroutines) // Semaphore to limit concurrent goroutines

	for start := 0; start < rows; start += chunkSize {
		end := start + chunkSize
		if end > rows {
			end = rows
		}

		sem <- struct{}{}
		wg.Add(1)
		go func(start, end int) {
			defer func() { <-sem }() // Release slot in the semaphore
			writeChunk(writer, start, end, &wg, &mu)
		}(start, end)
	}

	wg.Wait()
	return nil
}

func main() {
	rows := 1000000
	chunkSize := 10000
	maxGoroutines := 4 // Limit the number of concurrent goroutines to 4

	start := time.Now()

	err := writeConcurrently("output_concurrent.csv", rows, chunkSize, maxGoroutines)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Concurrent write complete in %s.\n", elapsed)
}
