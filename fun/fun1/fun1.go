package main

import "fmt"

//函数内部可以直接使用参数和返回值变量，而且 return 不需要显式地返回 x、y
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(36))
	fmt.Println(split(17))
}
