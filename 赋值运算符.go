package main

func main() {

	var a int = 21
	var b int

	//赋值运算符
	b = a
	println(b)

	//简化赋值运算符: += -+ *= /= %=
	//b+=a : b=a+b b-=a
	//b += a
	b %= a
	println(b)

}
