package main

/**
1). 实际上不能在基本类型，比如 int 上定义方法，但是简单重重命名一下，便可以了；
2). 可以在自己包内的类型上定义方法，不能对包含的其他包和基本类型定义方法；
3). 使用 float64 定义 MyFloat，可以将 float64 直接传给 MyFLoat，但是不能反过来直接传递，可以 float64(f) ；

**/
import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
