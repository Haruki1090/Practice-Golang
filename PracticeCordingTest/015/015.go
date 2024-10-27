// 問題 15: 2つの整数スライスが与えられたとき、両方のスライスに存在する共通の要素を持つ新しいスライスを返す関数をGoで実装してください。共通の要素がない場合は、空のスライスを返してください。スライス内での要素の順序は気にしなくても構いません。

// 例:

// 入力: []int{1, 2, 3, 4, 5} と []int{4, 5, 6, 7, 8} の場合、出力は []int{4, 5}
// 入力: []int{10, 20, 30} と []int{40, 50, 60} の場合、出力は []int{}
// ヒント:

// 1つ目のスライスの要素を map に登録し、2つ目のスライスをループで回しながら、map 内に共通する要素があるか確認すると便利です。

// 重複した共通要素が結果スライスに複数回追加される可能性があるため、少しだけ改善できます。たとえば、nums2 に重複する数が含まれていると、同じ共通要素が何度も result に追加されてしまいます。

// 改善ポイント
// map のもう一つの役割として、すでに result に追加された共通要素を記録しておくことで、重複する要素を防ぐことができます。

package main

import "fmt"

func commonElements(nums1, nums2 []int) []int {

	// 1つ目のスライスの要素を記録するmap
	seen := make(map[int]bool)

	// 共通の要素用のスライス
	result := []int{} // 空のスライスで初期化

	// 1つ目のスライスの要素をmapに追加
	for _, num := range nums1 {
		seen[num] = true
	}

	// 2つ目のスライスの要素をmap内で検索し、すでに追加されていないか確認
	uniqueSeen := make(map[int]bool) // 重複を防ぐためのmap
	for _, num := range nums2 {
		// 1つ目のスライスに存在し、かつ、まだ追加されていない場合 result に追加
		if seen[num] && !uniqueSeen[num] {
			result = append(result, num)
			uniqueSeen[num] = true // 追加した要素を記録
		}
	}

	return result
}

func main() {
	fmt.Println(commonElements([]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8})) // [4 5]
	fmt.Println(commonElements([]int{10, 20, 30}, []int{40, 50, 60}))       // []
	fmt.Println(commonElements([]int{1, 1, 2, 3, 4, 5}, []int{1, 1, 1, 1, 2, 2}))
}

// 変更点の説明
// uniqueSeen mapの追加
// uniqueSeen という map を追加し、共通要素がすでに result に追加されているかを追跡します。
// uniqueSeen[num] が false の場合のみ result に追加し、その後 uniqueSeen[num] = true として記録します。
// 重複を防ぐロジックの追加
// if seen[num] && !uniqueSeen[num] の条件を使用して、result に同じ要素が追加されないようにしています。
