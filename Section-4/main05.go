package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", fl64)

	fl := 10.5
	i3 := int(fl)
	fmt.Println(i3)

	var s string = "100"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	i4, _ := strconv.Atoi(s)
	fmt.Println(i4)
	fmt.Printf("%T\n", i4)

	var i5 int = 200
	s2 := strconv.Itoa(i5)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	var h string = "Hello, World!"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
