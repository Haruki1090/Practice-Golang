// 問題 5: 引数として2つの文字列を取り、その2つの文字列を連結して返す関数をGoで実装してください。

// 例:

// 入力: "Hello" と "World" の場合、出力は "HelloWorld"
// 入力: "Go" と "Lang" の場合、出力は "GoLang"

package main

import "fmt"

func concat(a, b string) string {
	return a + b
}

func main() {
	fmt.Println(concat("Hello", "World"))
	fmt.Println(concat("Go", "Lang"))
}
