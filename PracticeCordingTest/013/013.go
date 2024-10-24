// 問題 13: 整数のスライスが与えられたとき、そのスライス内に存在する重複する要素を削除し、ユニークな要素のみを含む新しいスライスを返す関数をGoで実装してください。

// 例:

// 入力: []int{1, 2, 2, 3, 4, 4, 5} の場合、出力は []int{1, 2, 3, 4, 5}
// 入力: []int{10, 20, 20, 30, 30, 30} の場合、出力は []int{10, 20, 30}
// ヒント:

// map を使うと、要素の重複を防ぐことができます。

package main

import "fmt"

func removeDuplicates(nums []int) []int {

	// 要素を記録するmap
	seen := make(map[int]bool)

	// 重複のない結果用のスライス
	result := []int{}

	for _, num := range nums {
		// 要素がmapに存在しない場合は、mapに追加し、結果用のスライスに追加
		if !seen[num] {
			result = append(result, num) // 重複していない要素を追加
			seen[num] = true             // 追加した要素をmapに追加
		}
	}

	return result
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 4, 4, 5}))    // [1 2 3 4 5]
	fmt.Println(removeDuplicates([]int{10, 20, 20, 30, 30, 30})) // [10 20 30]
}
