// 問題 9: 整数のスライスを引数に取り、スライスの中で最も頻繁に現れる要素を返す関数をGoで実装してください。同じ頻度の要素が複数ある場合は、その中のどれか一つを返してください。

// 例:

// 入力: [1, 2, 2, 3, 3, 3, 4] の場合、出力は 3
// 入力: [5, 5, 6, 6, 7, 7, 7, 6] の場合、出力は 6 か 7 （どちらでも正解）

package main

import "fmt"

func mostFrequent(num_slice []int) int {
	// 出現回数を保存するmapを作成
	count_map := make(map[int]int) // key: int, value: int

	// スライス内の要素の出現回数をカウント
	for _, num := range num_slice {
		count_map[num]++ // 既にkeyが存在していればvalueを+1、存在していなければ新たにkeyを作成しvalueを1にする
	}

	// 出現回数が最大の要素を探す
	most_frequent_num := num_slice[0] // 最も頻繁に出現する要素
	max_count := 0                    // 最も頻繁に出現する要素の出現回数

	// mapのkeyとvalueを取り出して、最も頻繁に出現する要素を探す
	for key, value := range count_map {

		// 出現回数が最大の要素を更新
		if value > max_count {
			most_frequent_num = key // 最も頻繁に出現する要素を更新
			max_count = value       // 最も頻繁に出現する要素の出現回数を更新
		}

	}

	return most_frequent_num
}

func main() {
	fmt.Println(mostFrequent([]int{1, 2, 2, 3, 3, 3, 4}))    // 3
	fmt.Println(mostFrequent([]int{5, 5, 6, 6, 7, 7, 7, 6})) // 6 or 7
}
