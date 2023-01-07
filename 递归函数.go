package main

func main() {
	println(getusm(2))
}

func getusm(num int) int {
	if num == 1 {
		return 1
	}
	return getusm(num-1) + num
}
