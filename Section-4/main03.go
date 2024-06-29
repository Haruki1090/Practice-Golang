package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 配列の要素数を省略することもできる
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)
	fmt.Println(arr4[1])
}

// Goでは、配列の要素数を指定して宣言する
// 配列の要素数は固定で、要素数を変更することはできない
