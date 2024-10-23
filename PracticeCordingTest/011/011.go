// 問題 11: 整数 n を引数に取り、1 から n までの数の中で素数だけを出力する関数をGoで実装してください。

// 例:

// 入力: 10 の場合、出力は 2, 3, 5, 7
// 入力: 20 の場合、出力は 2, 3, 5, 7, 11, 13, 17, 19

package main

import "fmt"

// 素数かどうかを判定
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 1からnまでの数の中で素数だけを出力
func printPrimes(n int) {
	for i := 2; i <= n; i++ {
		// 素数かどうかを判定
		// falseの場合は何もしない
		// trueの場合はその数を出力
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	printPrimes(10) // 2 3 5 7
	printPrimes(20) // 2 3 5 7 11 13 17 19
}
