package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()

	if id%2 == 0 { // IDが偶数ならエラーを出す
		errors <- fmt.Errorf("Worker %d failed", id)
		return
	}

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	errorChan := make(chan error, 10)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg, errorChan)
	}

	wg.Wait()
	close(errorChan)

	// エラーを出力
	for err := range errorChan {
		fmt.Println("Error:", err)
	}

	fmt.Println("All workers completed.")
}
