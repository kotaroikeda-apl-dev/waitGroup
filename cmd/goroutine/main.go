package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go hello()                  // 並列で実行
	time.Sleep(1 * time.Second) // メイン関数が終了しないようにする
}
