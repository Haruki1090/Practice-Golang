// 問題 2:
// 2つの整数を引数に取り、その合計を返す関数 add をGoで実装してください。次の例を参考にしてください。

// 例:
// a = 3, b = 8 のとき、a + b = 11 になります。
// a = 5, b = 7 のとき、a + b = 12 になります。

// 関数 add はこのような形で動作することが期待されます。

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(3, 8), add(5, 7))
}
