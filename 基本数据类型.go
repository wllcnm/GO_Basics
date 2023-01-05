package main

import "fmt"

func main() {
	//var 变量名 数据类型
	/*
		1、fmt.Println相比println的兼容性更好，支持更多形式的类型打印
		2、println不支持复杂类型和自定义类型打印
	*/

	//布尔型
	//bool : true false
	//bool默认值: false
	var isFlag bool = true
	var isFlag2 bool = false
	println(isFlag)

	fmt.Printf("%T,%t\n", isFlag, isFlag)
	fmt.Printf("%T,%t\n", isFlag2, isFlag2)

	//数字型

	//定义整型
	//byte uint8(0~255)
	//int 默认uint64
	var num int = 100
	//var age byte = 255
	fmt.Printf("%T,%d \n", num, num)

	//定义浮点型,
	//使用printf默认打印六位小数
	//使用float32会导致精度丢失,尽量使用float64
	var money1 float32 = -123.0000901
	var money2 float64 = -123.0000901
	fmt.Printf("%T,%f \n", money1, money1)
	fmt.Printf("%T,%f \n", money2, money2)
	fmt.Println(money1, money2)

	//字符串类型
	var str string
	str = "hello"
	fmt.Printf("%T,%s \n", str, str)

	//单引号并且是单字符,表示整型
	V1 := 'A' //65
	V2 := "A" //A
	println(V1, V2)

	//ASCELLA编码
	//中文编码:GBK
	//全世界编码表:Unicode

	//字符串连接,使用+号
	println("hello" + "world")

	//转义字符 \
	println("hello\"world")
	println("hello\nworld") // \n:换行
	println("hello\tworld") // \t:间隔
}
