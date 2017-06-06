package main

import "fmt"

func ExFunc(n int) func() {
	sum := n
	a := func() {
		fmt.Println(sum + 1)
	}
	return a
}
func main() {
	myFunc := ExFunc(10)
	myFunc() //这里输出11
	myAnotherFunc := ExFunc(20)
	myAnotherFunc() //这里输出21
	myFunc()        //这里输出11
	myAnotherFunc() //这里输出21

}
