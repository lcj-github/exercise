package main

import "fmt"

//make 函数可以用于构造数组，前面也构造过 map； new 用来构造结构体
func main() {
	a := make([]int, 5)
	printSlice("a", a) //a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b) //b len=0 cap=5 []
	c := b[:2]
	printSlice("c", c) //c len=2 cap=5 [0 0]
	d := c[2:5]
	printSlice("d", d) //d len=3 cap=3 [0 0 0]

	a[1] = 2
	printSlice("a", a) // a len=5 cap=5 [0 2 0 0 0]

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
