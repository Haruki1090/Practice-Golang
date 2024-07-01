package main

// 関数
// 関数を引数に取る関数

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		println("I'm a function")
	})
}
