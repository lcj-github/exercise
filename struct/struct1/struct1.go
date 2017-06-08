package main

import "fmt"

//1). 使用 type 和 struct 定义结构体  2). 结构体中类型声明和普通的一致，先写名字，再写类型
type Vertex struct {
	X int
	Y int
}

//构造一个结构体值，或是对象，或叫实例，直接名字加参数值即可
func main() {
	fmt.Println(Vertex{1, 2})

	//直接 v.X 就可以访问，既可以直接使用该值，也可以直接给该值赋值
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
}
