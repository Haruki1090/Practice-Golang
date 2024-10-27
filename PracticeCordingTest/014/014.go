// 問題 14: 整数のスライスが与えられたとき、そのスライス内で 最も頻繁に出現する要素 を返す関数をGoで実装してください。同じ回数出現する要素が複数ある場合は、そのうちのどれか一つを返して構いません。

// 例:

// 入力: []int{1, 2, 2, 3, 3, 3, 4} の場合、出力は 3
// 入力: []int{5, 5, 6, 6, 7, 7, 7, 6} の場合、出力は 6 か 7 （どちらでも正解）
// ヒント:

// map を使って各要素の出現回数をカウントすると便利です。

// ポイント
// mostFrequent と maxCount の初期値設定
// 現在のコードでは、mostFrequent と maxCount の初期値が 0 に設定されています。
// 入力のスライスに負の数が含まれる場合、誤って 0 が最頻出要素として選ばれる可能性があるので、mostFrequent は nums の最初の要素に、maxCount は 0 に初期化すると良いでしょう。

package main

import "fmt"

func mostFrequent(nums []int) int {
	// 要素の出現回数を記録するmap
	count := make(map[int]int)

	// 最も頻繁に出現する要素
	mostFrequent := nums[0] // 最初の要素で初期化 (負の数が含まれる場合に備えて)

	// 最も頻繁に出現する要素の出現回数
	maxCount := 0

	for _, num := range nums {
		// 出現数をカウント
		count[num]++

		// 最も頻繁に出現する要素を更新
		if count[num] > maxCount {
			mostFrequent = num
			maxCount = count[num]
		}
	}

	return mostFrequent
}

func main() {
	fmt.Println(mostFrequent([]int{1, 2, 2, 3, 3, 3, 4}))    // 3
	fmt.Println(mostFrequent([]int{5, 5, 6, 6, 7, 7, 7, 6})) // 6 or 7
}
