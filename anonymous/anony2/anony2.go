package main

import (
	"fmt"
)

func main() {

	//无参数直接加括号
	func() int {
		var i int = 5
		fmt.Printf("func 1\n")
		return i
	}() //func 1

	//有参数，在括号里加参数
	func(arge int) {
		fmt.Printf("func %d\n", arge)
	}(2) //func 2

	//也可以先赋给一个变量再调用
	a := func() int {
		fmt.Printf("func 3\n")
		return 5
	}
	a() //func 3

}
