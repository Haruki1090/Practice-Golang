// 問題 4: 整数のスライスを引数に取り、その中の偶数のみを含む新しいスライスを返す関数をGoで実装してください。

// 例:

// 入力: [1, 2, 3, 4, 5, 6] のとき、出力は [2, 4, 6]
// 入力: [7, 11, 15, 20] のとき、出力は [20]

package main

import "fmt"

func even(slice []int) []int {

	var evenSlice []int

	// スライス内の全ての要素を順番にループする
	for _, v := range slice {
		// v にはスライスの各要素が順に入る.
		// インデックスは使わないので _ で無視しています。
		// `range slice` でスライスの全要素を取り出すことができるので、わざわざインデックスを使う必要はありません。
		// Goは未使用の変数を許可しないので、_ で無視しています。（未使用なのに書くと警告出る）

		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

func main() {
	fmt.Println(even([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(even([]int{7, 11, 15, 20}))
}
