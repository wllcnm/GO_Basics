package main

import "fmt"

// 全局变量
var globalVariable string = "全局变量"

func main() {
	//定义字符串类型变量
	var name string = "jojo"
	//定义整型类型变量
	var age int = 18
	//定义多个变量
	var (
		add    string
		phone  string
		gender string
	)
	//对变量重新赋值
	name = "承太郎"
	add = "广西"
	phone = "18777355162"
	gender = "男"

	// := 自动推导
	email := "1712@123.com"

	println(name)

	println(age)

	println(add, phone, gender, email)

	//打印变量的类型,通过printf
	//email变量的类型为string
	fmt.Printf("email变量的类型为%T \n", email)

	//打印变量的内存存地址,通过printf
	//email的内存地址0xc00004e260
	fmt.Printf("email的内存地址%p \n", &email)

	//1.a=temp
	//2.a=b
	//3.b=temp

	//go语言变量交换,通过a,b=b,a
	a := 100
	b := 200
	a, b = b, a
	fmt.Printf("a的值为%d,b的值为%d", a, b)

	//匿名变量,用下划线表示 _ ,表示废弃的
	//使用两个变量接受一个函数的返回值,如果不想要某一个值可以用下划线表示
	c, _ := test()
	println(c)

	//局部变量,在函数体内作用
	localVariable := "局部变量"
	//定义了全局变量也可以在函数内重新定义,就近原则,使用函数内定义的变量
	globalVariable := "新的全局变量"
	fmt.Printf("%s,%s", globalVariable, localVariable)

}

func test() (int, int) {
	return 100, 200
}
