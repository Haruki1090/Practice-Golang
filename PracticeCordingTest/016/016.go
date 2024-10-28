// 問題 16: 文字列が与えられたとき、その文字列が回文（前から読んでも後ろから読んでも同じ）かどうかを判定する関数をGoで実装してください。
// 大文字と小文字は区別せず、スペースや記号も無視して判定するものとします。

// 例:

// 入力: "A man, a plan, a canal, Panama!" の場合、出力は true（回文）
// 入力: "Hello, World!" の場合、出力は false（回文ではない）
// ヒント:

// 文字列を正規化（すべて小文字にし、アルファベット以外の文字を除去）する処理が必要です。
// その後、両端から文字を比較していく方法で判定すると良いでしょう。

package main

import (
	"fmt"
	"strings" // 文字列操作を行うための標準パッケージ
	"unicode" // Unicode文字の分類や変換を行うための標準パッケージ
)

// 文字列が回文かどうかを判定する関数
func isPalindrome(s string) bool {
	// 1. 文字列を小文字に変換し、アルファベット以外の文字を除去
	var cleaned strings.Builder // 文字列を構築するための構造体
	for _, r := range s {
		if unicode.IsLetter(r) {
			cleaned.WriteRune(unicode.ToLower(r)) // 小文字に変換して追加
		}
	}

	// 2. 左右から文字を比較して回文かどうかを確認する
	cleanedStr := cleaned.String()      // 文字列に変換. cleaned は strings.Builder 型なので、string 型に変換する
	left, right := 0, len(cleanedStr)-1 // 左端と右端のインデックス. 0 から始まるので、右端は長さ-1
	for left < right {
		if cleanedStr[left] != cleanedStr[right] {
			return false // 左右の文字が異なる場合は回文ではない
		}
		left++  // 左端を右に移動
		right-- // 右端を左に移動
	}

	return true // 左右の文字が全て一致した場合は回文！
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal, Panama!")) // true
	fmt.Println(isPalindrome("Hello, World!"))                   // false
}
