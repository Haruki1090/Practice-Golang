// 問題 12: 整数のスライスが与えられたとき、そのスライスの要素を逆順に並び替える関数をGoで実装してください。

// 例:

// 入力: []int{1, 2, 3, 4, 5} の場合、出力は []int{5, 4, 3, 2, 1}
// 入力: []int{10, 20, 30} の場合、出力は []int{30, 20, 10}

package main

import "fmt"

func reverce(nums []int) []int {
	start := 0
	end := len(nums) - 1 // 末尾のインデックス. 0から始まるので-1する

	for start < end {
		// nums[start] と nums[end] を交換
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}

func main() {
	fmt.Println(reverce([]int{1, 2, 3, 4, 5})) // [5 4 3 2 1]
	fmt.Println(reverce([]int{10, 20, 30}))    // [30 20 10]
}
