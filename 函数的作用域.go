package main

var num int = 100

func main() {

	//如果定义了全局变量和局部变量,遵循就近原则
	println(num)
	f1()
	f2()
}

func f1() {
	num := 20
	println(num)
}
func f2() {
	num := 30
	println(num)
}
