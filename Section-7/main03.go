package main

// 関数
// 関数をかえす関数

func ReturnFunc() func() {
	return func() {
		println("I'm a function")
	}
}

func main() {
	f := ReturnFunc()
	f()
}
