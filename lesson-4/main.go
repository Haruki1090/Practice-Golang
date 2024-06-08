package main

import "fmt"

// 定数
const Pi = 3.14

// 定数のグループ
const (
	URL      = "http://example.com"
	SiteName = "Golang"
)

// 定数の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

// iota
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}
