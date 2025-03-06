package main

import (
	"fmt"
	"sync"
	"time"
)

const maxWorkers = 3 // 最大並列数

func worker(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for job := range ch {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(2 * time.Second) // 2秒待つ
	}
}

func main() {
	var wg sync.WaitGroup
	jobQueue := make(chan int, 10)

	// Worker を3つ起動
	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobQueue)
	}

	// ジョブを追加
	for j := 1; j <= 10; j++ {
		jobQueue <- j
	}
	close(jobQueue) // チャンネルを閉じる

	wg.Wait() // すべての Worker の終了を待つ
	fmt.Println("All jobs completed.")
}
