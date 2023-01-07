package main

func main() {

	//if语句
	score := 90
	if score >= 90 {
		println("成绩A")
	} else if score < 90 && score >= 80 {
		println("成绩B")
	} else if score < 80 && score >= 70 {
		println("成绩C")
	} else if score < 70 && score >= 60 {
		println("成绩D")
	} else {
		println("不及格")
	}

	//if嵌套判断
	//var a int
	//var b int
	//var pwd int = 2000300421
	//println("请输入密码")
	//fmt.Scanln(&a)
	//if a == pwd {
	//	println("请再次输入密码")
	//	fmt.Scanln(&b)
	//	if b == pwd {
	//		println("登录成功")
	//	} else {
	//		println("第二次输入密码错误,登陆失败")
	//	}
	//} else {
	//	println("密码错误")
	//}

	//switch语句
	switch score {
	case 90:
		println("A")
	case 80:
		println("B")
	case 70:
		println("C")
	case 60:
		println("D")
	default:
		println("没有匹配值")
	}

	//switch的默认条件 bool=true
	switch {
	case true:
		println("switch的默认条件 bool=true")
	}

	//fallthough
	switch score {
	case 90:
		println("A")
		fallthrough
	case 80:
		if score == 90 {
			break //break可强制退出
		}
		println("B")
		fallthrough
	case 70:
		println("C")
		fallthrough
	case 60:
		println("D")
	default:
		println("没有匹配值")
	}

	//for循环
	for i := 1; i <= 10; i++ {
		println(i)
	}
	//计算 1+2+3+4+.......
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	println("总和为", sum)
	//for如果没有参数就是无限循环,想当于while
	//for循环可以有一个或两个或三个
	//如果在外面定义了参数,for的第一个参数可以不写
	i := 1
	for i < 10 {
		println(i)
		i += 2
	}

	//打印5X5方阵
	for i := 0; i < 5; i++ {
		println()
		for i := 0; i < 5; i++ {
			print("*")
		}
	}

	//break:结束本次循环
	println("结束本次循环")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}
	//continue:结束单次循环
	println("结束单次循环")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		println(i)
	}

}
