package main

import "fmt"

func main() {
	a := 3
	b := 5.0

	//整型和浮点数之间进行转换
	d := float64(a)
	e := int(b)
	fmt.Printf("%T %d \n", a, a)
	fmt.Printf("%T %f \n", b, b)
	fmt.Printf("%T %f \n", d, d)
	fmt.Printf("%T %d \n", e, e)
}
