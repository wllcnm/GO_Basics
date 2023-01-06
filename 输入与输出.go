package main

import "fmt"

func main() {

	var x int
	var y float64

	println("请输入一个整数和浮点数")
	//&:取地址符 控制台输入的语法和c语言差不多
	fmt.Scanln(&x, &y)
	println("x:", x)
	fmt.Println("y:", y)

}
