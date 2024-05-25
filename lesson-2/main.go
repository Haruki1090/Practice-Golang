package main

import "fmt"

// 変数宣言

// i5 := 500 // エラー
var i5 int = 500 // これが正しい

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 明示的な型宣言
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	i = 150
	fmt.Println(i)

	// 暗黙的な型宣言
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450 // 再代入
	fmt.Println(i4)

	// i4 = "Hello" // 型が違うのでエラー

	fmt.Println(i5)

	outer() // outer関数内で宣言された変数はmain関数内では使えない
}
