package main

import (
	"fmt"
)

type Animal interface {
	move()
}
type Human struct {
	i int
}

func (r Human) move() { // 注意
	fmt.Println("人类行走")
	r.i++
}

type Bird struct {
	i int
}

//receiver为指针，所以三次调用后，输出值累加
func (r *Bird) move() {
	fmt.Println("鸟类行走")
	r.i++
}
func moveTest1(animal Animal) {
	animal.move()
}

// 虽然这个函数的定义没有错误，但实际上完全用不上！
// 因为指向接口的指针，在接口实现类中，是无法转换的
func moveTest2(animal *Animal) {
	(*animal).move()
}

func main() {
	h1 := Human{0}
	moveTest1(h1)
	moveTest1(h1)
	moveTest1(h1)
	fmt.Println(h1.i) //0
	fmt.Println("----------------")

	h2 := &Human{0}
	moveTest1(h2)
	moveTest1(h2)
	moveTest1(h2)
	fmt.Println(h2.i) //0
	fmt.Println("----------------")

	b2 := &Bird{0}
	moveTest1(b2)
	moveTest1(b2)
	moveTest1(b2)
	fmt.Println(b2.i) //3
	fmt.Println("----------------")

	//  h3 := Human{0}
	//  moveTest2(h3)
	//  moveTest2(h3)
	//  moveTest2(h3)
	//  fmt.Println(h3.i)
	//
	//  fmt.Println("-------如上报错---------")

	//  b1 := Bird{0}
	//  moveTest1(b1)
	//  moveTest1(b1)
	//  moveTest1(b1)
	//  fmt.Println(b1.i)
	//
	//  fmt.Println("--------如上报错--------")

	//  b3 := &Bird{0}
	//  moveTest2(b3)
	//  moveTest2(b3)
	//  moveTest2(b3)
	//  fmt.Println(b3.i)
	//
	//  fmt.Println("-------如上报错---------")
}
