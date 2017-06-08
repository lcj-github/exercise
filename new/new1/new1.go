package main

import "fmt"

type Vertex struct {
	X, Y int
}

//1). new 构造结构体，返回指针 2). new 的同时貌似不能初始化，new 是个函数，不像 c++ 中是个运算符；
//&{0 0}  &{11 9}
func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
