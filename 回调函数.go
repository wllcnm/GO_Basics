package main

import "fmt"

func main() {

	r2 := Add(1, 2)
	println(r2)

	//回调函数做加法
	r3 := oper(1, 2, Add)
	println(r3)

	//回调函数做减法
	r4 := oper(1, 2, sub)
	println(r4)

	//回调函数做乘法
	r5 := oper(3, 2, mul)
	println(r5)

	//回调加匿名函数做除法
	r6 := oper1(3, 2, func(i float64, i2 float64) float64 {
		return i / i2
	})
	fmt.Printf("%f", r6)
}

// 回调函数,可以接受一个函数作为参数
func oper(a, b int, add func(int, int) int) int {
	return add(a, b)
}
func oper1(a, b float64, add func(float64, float64) float64) float64 {
	return add(a, b)
}

func Add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) float64 {
	return float64(a / b)
}
