# Go の Goroutine を使ったサンプル

## 概要

このプロジェクトは、Go の`goroutine`を使って並列処理を実行する方法を示します。`go`キーワードを使った並列処理の基本を理解し、適切に制御する方法を学びます。

## `goroutine`とは？

`goroutine`は、Go 言語が提供する軽量なスレッドのような仕組みです。通常の関数呼び出しの前に`go`キーワードを付けることで、新しい`goroutine`として並列実行できます。

### 学んだポイント

- **goroutine の作成**: `go 関数名()` を使って並列処理を開始。
- **並列処理の非同期実行**: goroutine はメイン処理とは独立して実行される。
- **メイン関数の終了による goroutine の影響**: メイン関数が終了すると、実行中の goroutine も終了してしまう。
- **`time.Sleep()` の活用**: goroutine が終了する前にメイン関数を維持するための一時的な手段。
- **`sync.WaitGroup` の活用**: `goroutine`の完了を適切に待つための手段。
- **`defer` の活用**: `WaitGroup`を適切に管理し、処理の終了を保証。

### 実行方法

```sh
go run main.go  # goroutine を使った並列処理の実行
```

### 学習ポイント

1. **`go`キーワードを使うことで関数を並列実行できる**
2. **メイン関数が終了すると、goroutine も強制終了する**
3. **`time.Sleep()` は goroutine を維持するための一時的な方法だが、適切な方法ではない**
4. **`sync.WaitGroup` を使うことで、goroutine が完了するまでメイン関数の終了を待つことができる**
5. **`defer wg.Done()` を使うことで goroutine の終了時にカウントを適切に減らすことができる**

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
