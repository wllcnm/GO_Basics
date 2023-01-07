package main

func main() {
	//使用defer关键字会使该函数延迟到最后执行
	//如果有多个defer函数,会逆序执行
	defer f("1")
	defer f("3")
	f("2")
}
func f(msg string) {
	println(msg)
}
