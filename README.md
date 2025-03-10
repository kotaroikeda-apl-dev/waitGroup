# Go の Goroutine を使ったサンプル

## 概要

このプロジェクトは、Go の `goroutine` を使って並列処理を実行する方法を示します。`go` キーワードを使った並列処理の基本を理解し、適切に制御する方法を学びます。

## `goroutine` とは？

`goroutine` は、Go 言語が提供する軽量なスレッドのような仕組みです。通常の関数呼び出しの前に `go` キーワードを付けることで、新しい `goroutine` として並列実行できます。

### **実行方法**

```sh
go run cmd/goroutine/main.go  # goroutine を使った並列処理の実行
go run cmd/waitGroup/main.go  # sync.WaitGroupの活用
go run cmd/limit/main.go  # ループ数を変数で管理
go run cmd/error/main.go  # エラー処理
go run cmd/context/main.go  # タイムアウト制御
```

## **学習ポイント**

1. **`go` キーワードを使うことで関数を並列実行できる**
2. **メイン関数が終了すると、goroutine も強制終了する**
3. **`time.Sleep()` は goroutine を維持するための一時的な方法だが、適切な方法ではない**
4. **`sync.WaitGroup` を使うことで、goroutine が完了するまでメイン関数の終了を待つことができる**
5. **`defer wg.Done()` を使うことで goroutine の終了時にカウントを適切に減らすことができる**
6. **チャネル (`chan`) を使うことで、goroutine 同士が安全にデータを受け渡せる**
7. **`close(jobQueue)` を適切に行うことで、goroutine の終了を制御できる**
8. **バッファ付きチャネルの仕組みを理解し、ブロッキングの挙動を学ぶ**
9. **ワーカーを複数並列に起動し、ジョブを分配する方法を学ぶ**
10. **ワーカー内でエラーが発生した際に `errorChan` を通じてエラーを収集し、後から処理する方法を学ぶ**
11. **`close(errorChan)` を適切に行うことで、エラーチャネルの `range` ループが正常に終了する仕組みを理解する**
12. **`context.WithTimeout` を用いたタイムアウト制御を学ぶ**
13. **`select` を用いた goroutine のキャンセル処理を学ぶ**
14. **goroutine 内で `select` が 1 回しか実行されないことを理解する**
15. **`time.After(N * time.Second)` が `ctx.Done()` より先に実行される場合、キャンセル処理が発生しないことを学ぶ**
16. **goroutine 内で `select` をループさせることで、継続的に `ctx.Done()` を監視する手法を理解する**

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
