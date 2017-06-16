package main

import (
	"fmt"
	"math"
)

/**
1). 接口是一种类型，里面装的是一系列方法声明，也就是多个方法声明；
2). 一个接口值可以用任一个实现了该接口(实现其方法)的实例/对象赋值；
3). 实现接口的如果是指针类型 *T，那么，只能使用 &T 变量赋值给接口，而不能是 T 。
**/
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	a = &v // a *Vertex implements Abser
	//a = v  // a Vertex, does NOT
	// implement Abser

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
