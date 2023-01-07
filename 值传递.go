package main

import "fmt"

func main() {
	//数组 [个数]类型{具体值}
	args := [3]int{1, 2, 3}
	fmt.Println(args)
	update(args)
	//值传递进函数体只是进行了拷贝,不会影响具体的值,类似于c语言
	fmt.Println(args)
}
func update(args2 [3]int) [3]int {
	args2[0] = 100
	return args2
}
