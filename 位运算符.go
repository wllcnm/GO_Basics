package main

import "fmt"

func main() {

	//位运算:二进制上的0 false 1 true
	//60 0011 1100
	//13 0000 1101
	//60&13 0000 1100 => 12
	//60||13 0011 1101 => 61
	println(60 & 13)
	println(60 | 13)

	var a int = 13
	var b int = 60
	//%b表示数值的二进制 binary
	fmt.Printf("二进制值%b \n", a&b)

	//按位异或,相同取0,不同取1
	//60^13 0011 0001
	fmt.Printf("%b \n", a^b)

	//按位左移或右移
	//0000 1101 << 2 =>13左移两位:高位舍弃低位补零
	//0000 1101 >> 2 =>13右移两位:低位舍弃高位补零
	fmt.Printf("%b \n", a<<2)
	fmt.Printf("%b \n", a>>2)
}
