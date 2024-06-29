package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println(1 + 2)
	fmt.Println(5 - 2)
	fmt.Println(6 * 6)
	fmt.Println(8 / 2)
	fmt.Println(9 % 4)
	fmt.Println("Hello" + "・ " + "Go")

	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "Hello"
	s += " Go"
	fmt.Println(s)
}
