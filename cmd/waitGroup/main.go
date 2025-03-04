package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ゴルーチンが終わったらカウントを減らす
	fmt.Printf("Worker %d starting\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // ゴルーチンのカウントを追加
		go worker(i, &wg)
	}

	wg.Wait() // すべてのゴルーチンが終了するのを待つ
	fmt.Println("All workers finished.")
}
