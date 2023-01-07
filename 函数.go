package main

// main函数是入口
func main() {

	i := add(1, 3)
	println(i)

	printfinfo()
	printfinfo1("printfinfo1")
	printinfo2("123", 20)
	println(swap("cxk", "jinitaimei"))
	println("最大值为", max(20, 10))
	println(sum("haha", 1, 2, 3, 4, 5, 56, 2))
}

func add(a, b int) int {
	return a + b
}

// 无参无返回值的函数
func printfinfo() {
	println("printinfo")
}

// 一个参数无返回值
func printfinfo1(msg string) {
	println(msg)
}

// 两个参数且有返回值
func printinfo2(msg string, num int) (string, int) {
	return msg, num
}

// func swap 交换函数
func swap(x, y string) (string, string) {
	x, y = y, x
	return x, y
}

// func max 求最大值函数
// 这里的x和y为形参,传入的为实参
func max(x, y int) int {
	if x == y {
		return x
	}
	if x > y {
		return x
	} else {
		return y
	}
}

// 可变参数
// ps:可变参数也可以和固定参数组合
func sum(msg string, args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
