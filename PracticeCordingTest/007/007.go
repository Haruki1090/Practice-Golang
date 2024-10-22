// 問題 7: 整数のスライスを引数に取り、その中の要素の合計を返す関数をGoで実装してください。

// 例:

// 入力: [1, 2, 3, 4, 5] のとき、出力は 15
// 入力: [10, -2, 3] のとき、出力は 11

package main

import "fmt"

func sum(num_slice []int) int {
	sum := 0
	for _, v := range num_slice {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(sum([]int{10, -2, 3}))
}
