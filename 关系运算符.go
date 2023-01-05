package main

import "fmt"

func main() {
	var a int = 10
	var b int = 11

	//关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	if a < b {
		println("a比b大")
	} else {
		println("a比b小")
	}
	
}
