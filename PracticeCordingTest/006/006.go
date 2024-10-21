// 問題 6: 1から100までの整数を出力するプログラムをGoで書いてください。ただし、3の倍数のときは「Fizz」、5の倍数のときは「Buzz」、3と5の両方の倍数のときは「FizzBuzz」と出力するようにしてください。

// 例:

// 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, ...

package main

import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

// fmt.Printについて
//  fmt.Print は Go の標準ライブラリにある関数で、指定した内容を 改行なし で出力
//  fmt.Println は fmt.Print と違い、改行を含む出力を行う関数

// fmt.Print(i, " ")について
// fmt.Print(i, " ") は、i と " "（半角スペース）を出力するという意味
// このようにすることで、出力される数字がスペースで区切られるようになる

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
