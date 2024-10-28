// 問題 18: 整数のスライスが与えられたとき、そのスライスの中で最も頻繁に出現する要素の出現回数を返す関数をGoで実装してください。

// 例:

// 入力: []int{1, 2, 2, 3, 3, 3, 4} の場合、出力は 3 （数値 3 が3回出現する）
// 入力: []int{5, 5, 6, 6, 7, 7, 7, 6} の場合、出力は 3 （数値 7 が3回出現する）
// 入力: []int{1, 1, 1, 2, 2} の場合、出力は 3 （数値 1 が3回出現する）
// ヒント:

// map を使って各要素の出現回数をカウントし、その中で最大の出現回数を見つけると便利です。

package main

import "fmt"

// 出現回数をカウントする関数
func countOccurrences(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	return counts
}

// 出現回数のmapから最大の出現回数を返す関数
func maxCount(counts map[int]int) int {
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

// スライスの中で最も頻繁に出現する要素の出現回数を返す関数
func mostFrequentCount(nums []int) int {
	if len(nums) == 0 {
		return 0 // 空のスライスに対応
	}
	counts := countOccurrences(nums)
	return maxCount(counts)
}

func main() {
	// テスト
	fmt.Println(mostFrequentCount([]int{1, 2, 2, 3, 3, 3, 4}))    // 出力: 3
	fmt.Println(mostFrequentCount([]int{5, 5, 6, 6, 7, 7, 7, 6})) // 出力: 3
	fmt.Println(mostFrequentCount([]int{1, 1, 1, 2, 2}))          // 出力: 3
}

// 注意したポイント
// 1. 空のスライスに対応するためのエラーハンドリング
// 		スライスが空の場合（[]int{}）でも問題なく動作するが、より安全にするために、空のスライスの場合は 0 を返すようにしている。
// 2. mostFrequentCount 関数をさらに分割して、可読性を向上
//		関数の役割を明確にするため、出現回数をカウントする処理と、最大出現回数を探す処理を別の関数に分割した。
