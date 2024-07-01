package main

import "fmt"

// 関数
// クロージャー

func Later() func(string) string {
	var store string

	// 無名関数
	return func(next string) string {
		s := store
		store = next
		return s
	}
	// 関数の説明
	// 1. Later 関数が実行される
	// 2. store に空文字が代入される
	// 3. 無名関数が実行される
	// 4. store に next が代入される
	// 5. s に store の値が代入される
	// 6. s が返される
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("Is"))
	fmt.Println(f("Golang"))
}
