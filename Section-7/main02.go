package main

import "fmt"

// 関数
// 無名関数

func main() {
	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	// 即時関数
	i2 := func(x, y int) int {
		return x + y
	}(2, 3)
	fmt.Println(i2)
}
