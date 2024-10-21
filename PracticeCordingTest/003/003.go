// 問題 3:
// 整数のスライス（配列のようなもの）を引数に取り、その中の最大値を返す関数をGoで実装してください。

// 例:

// 入力: [1, 5, 3, 9, 2] のとき、出力は 9
// 入力: [7, 4, 0, -2] のとき、出力は 7

package main

import "fmt"

func max(slice []int) int {
	max := slice[0]

	for _, v := range slice {
		// スライス内の全ての要素を順に取り出すためのループです。
		// v にはスライスの各要素が代入されますが、インデックス（位置）は使わないので _ で無視しています。

		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(max([]int{1, 5, 3, 9, 2}))
	fmt.Println(max([]int{7, 4, 0, -2}))
}
