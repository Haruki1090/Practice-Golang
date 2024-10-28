// 問題 19: 与えられた文字列内で最も頻繁に出現する文字を見つけ、その出現回数を返す関数をGoで実装してください。空白や記号は無視し、大文字小文字も区別しないものとします。

// 例:

// 入力: "Hello, World!" の場合、出力は l: 3（文字 l が3回出現）
// 入力: "A man, a plan, a canal, Panama!" の場合、出力は a: 8（文字 a が8回出現）
// ヒント:

// map を使って、各文字の出現回数をカウントすると便利です。
// unicodeパッケージ を使って、文字がアルファベットかどうかを判定し、小文字に変換することができます。

package main

import (
	"fmt"
	"unicode"
)

// 最も頻繁に出現する文字とその出現回数を返す関数
func mostFrequentChar(s string) (rune, int) { // rune は Unicode のコードポイントを表す型
	// 各文字の出現回数を記録するmap
	counts := make(map[rune]int)
	// rune は Unicode のコードポイントを表す型

	// 文字列を走査してアルファベットの小文字のみをカウント
	for _, r := range s {
		if unicode.IsLetter(r) { // アルファベットの場合
			r = unicode.ToLower(r) // 小文字に変換
			counts[r]++            // 出現回数をカウント
		}
	}

	// 最も頻繁に出現する文字を探す
	var mostFrequent rune // 最も頻繁に出現する文字
	maxCount := 0         // 最大の出現回数
	for char, count := range counts {
		if count > maxCount {
			mostFrequent = char
			maxCount = count
		}
	}

	return mostFrequent, maxCount
}

func main() {
	// テスト
	char, count := mostFrequentChar("Hello, World!")
	fmt.Printf("%c: %d\n", char, count) // 出力例: l: 3

	char, count = mostFrequentChar("A man, a plan, a canal, Panama!")
	fmt.Printf("%c: %d\n", char, count) // 出力例: a: 8
}

// コードの解説
// 1. counts マップの初期化:
// ・counts という map[rune]int を作り、各文字の出現回数を記録します。

// 2. 文字列の走査とカウント:
// ・for _, r := range s で文字列 s の各文字を走査します。
// ・unicode.IsLetter(r) を使ってアルファベットだけを対象にし、さらに unicode.ToLower(r) で小文字に変換してカウントしています。

// 3. 最も頻繁に出現する文字の特定:
// ・counts マップ内の出現回数を走査し、maxCount よりも大きい値が見つかった場合、その文字と出現回数を更新します。
// ・最終的に mostFrequent に最頻出の文字が、maxCount にその回数が保存されます。
// 4. 結果の返却:
// ・mostFrequentChar 関数は最も頻繁に出現する文字（rune）とその出現回数（int）を返します。
