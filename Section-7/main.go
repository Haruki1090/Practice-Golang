package main

import "fmt"

// 関数

func plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	// rerurn のあとは省略できる
	return
}

func NoReturn() {
	fmt.Println("No Return")
}

func main() {
	i := plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 2)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	NoReturn()
}
