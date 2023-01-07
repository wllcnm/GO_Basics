package main

func main() {

	f7 := func() {
		println("我是匿名函数")
	}
	f7()

	func() {
		println("匿名函数自己调用自己")
	}()

	r1, r2 := func(a, b int) (int, int) {
		println(a, b)
		return a, b
	}(1, 2)
	println(r1, r2)
}
