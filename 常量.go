package main

func main() {
	//定义常量,常量是不可改变的,使用关键字const定义常量
	const url string = "jojo.lw123.top" //显示定义

	const url2 = "www.baidu.com" //隐式定义

	const A, B, C = 3.14, "123", "456" //定义了三个常量

	//iota:常量计数器
	//ps:1.如果常量没有定义会沿用上一个常量的值
	//2.iota计数器是自动递增的,无论有没有调用
	const (
		a = iota
		b
		c
		d = "haha"
		e
		f = 100
		g = iota
		h
	)

	const (
		j = iota
		k
	)
	//从打印结果可以看出,iota无论有没有调用都会自动递增,a~h对应0~7,且没有赋值的常量会沿用上一个常量的值
	println(a, b, c, d, e, f, g, h)

}
