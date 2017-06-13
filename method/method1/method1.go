package main

import (
	"fmt"
	"math"
)

/**
1). 所谓的“方法”(method) —— 就是特殊的函数，只不过与某个类型绑定起来了而已，比如 Abs 绑定到了 *Vertex 类型；
2). Abs 绑定到了 *Vertex 上之后，就可以使用 *Vertex 值去调用这个方法，比如 v.Abs() ；
3). Abs 绑定到了 *Vertex 上之后，函数体内部可以使用 *Vertex 可以访问到的东西，比如 v.X
**/
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	//fmt.Println((*v).Abs()) // ok,2
}
