// 問題 10: 文字列のスライスを引数に取り、それらをアルファベット順にソートして返す関数をGoで実装してください。

// 例:

// 入力: []string{"banana", "apple", "cherry"} の場合、出力は []string{"apple", "banana", "cherry"}
// 入力: []string{"dog", "cat", "bird"} の場合、出力は []string{"bird", "cat", "dog"}

package main

import (
	"fmt"
	"sort"
)

func sortStrings(str_slice []string) []string {
	// 文字列のスライスをアルファベット順にソート
	sort.Strings(str_slice)
	return str_slice
}

func main() {
	fmt.Println(sortStrings([]string{"banana", "apple", "cherry"})) // ["apple", "banana", "cherry"]
	fmt.Println(sortStrings([]string{"dog", "cat", "bird"}))        // ["bird", "cat", "dog"]
}
