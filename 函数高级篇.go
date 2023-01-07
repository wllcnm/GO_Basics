package main

import "fmt"

func main() {
	//函数不加括号,函数就是一个变量
	fmt.Printf("%T", f3) //函数的类型为func()

	//定义函数类型的变量
	var f4 func(int, int) int
	//将f3赋值给f4
	f4 = f3
	f4(1, 2)
	fmt.Printf("%T", f4)
}
func f3(a, b int) int {
	println(a, b)
	return 0
}
