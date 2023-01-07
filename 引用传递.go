package main

import "fmt"

func main() {

	//定义一个切片,切片:没有定义具体长度的数组
	args := []int{1, 2, 3, 4}
	fmt.Println(args)
	//引用传递,传入的是地址
	update2(args)
	fmt.Println(args)
}

func update2(args []int) {
	args[0] = 100
}
