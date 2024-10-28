// 問題 17: 整数 n が与えられたとき、n 以下のフィボナッチ数列をすべて出力する関数をGoで実装してください。
// フィボナッチ数列は、各項が前の2つの項の和で定義される数列で、最初の2項は 0 と 1 です。

// 例:

// 入力: n = 10 の場合、出力は [0, 1, 1, 2, 3, 5, 8]
// 入力: n = 20 の場合、出力は [0, 1, 1, 2, 3, 5, 8, 13]
// ヒント:

// フィボナッチ数列の次の項は、直前の2つの数を足したものです。
// スライスを使って、生成した数列を保存していくと便利です。

package main

import "fmt"

// n 以下のフィボナッチ数列を返す関数
func fibonacci(n int) []int {
	if n < 0 {
		return nil // 負の数の場合は空のスライスを返す
	}

	// 初期値として 0, 1 を設定
	result := []int{0, 1}

	// フィボナッチ数列を計算
	for {
		// 次の項を計算. 直前の2つの数を足す
		next := result[len(result)-1] + result[len(result)-2]
		// n を超える場合は終了
		if next > n {
			break // ループを抜ける
		}
		// スライスに追加
		result = append(result, next)
	}

	return result
}

func main() {
	fmt.Println(fibonacci(10))  // 出力: [0 1 1 2 3 5 8]
	fmt.Println(fibonacci(20))  // 出力: [0 1 1 2 3 5 8 13]
	fmt.Println(fibonacci(100)) // 出力: [0 1 1 2 3 5 8 13 21 34 55 89]
}
