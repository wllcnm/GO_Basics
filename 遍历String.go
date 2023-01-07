package main

import "fmt"

func main() {
	str := "hello"
	println(str)

	//获取字符串的长度 len
	println("str长度为", len(str))

	//获取指定的字节,单个字符打印会打印ASSIC码
	println("字节打印", str[0])
	//%c表示单个字符
	fmt.Printf("%c \n", str[0])
	//循环打印单个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	//for range循环
	//range循环:i表示下标,v表示下标表示的值
	for i, v := range str {
		print(i)
		fmt.Printf("%c \n", v)
	}

}
