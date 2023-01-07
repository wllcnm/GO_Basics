package main

import "fmt"

func main() {

	r1 := increment()
	fmt.Println(r1)

	v1 := r1()
	println(v1)
	println(r1())
	println(r1())
	println(r1())
	println(r1())

	//r2
	r2 := increment()
	println(r2())
	println(r2())
	println(r2())
	println(r2())
}

// 闭包结构:外层函数有内层函数,在该内层函数中,会操作外层函数的值
// 该外层函数的返回值就是内层函数的返回值
// 总结:闭包结构就是将函数作为返回值
func increment() func() int {
	//i:局部变量
	i := 0
	fun := func() int { //内层函数没有执行
		i++
		return i
	}
	return fun
}
