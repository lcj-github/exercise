package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//如果这个返回函数使用了原来函数内的变量值，那么这些变量的生命周期便延长了
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 4; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
