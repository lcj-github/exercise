package main

import "fmt"

//http://ilovers.sinaapp.com/node/33
type Vertex struct {
	X int
	Y int
}

//要访问指针 p 指向的结构体中某个元素 x，不需要显式地使用 * 运算，可以直接 p.x ；
func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 10
	(*q).Y = 20
	r := &q
	//(*r).Y = 4 // ok
	//(*(*r)).X = 100 //ok
	s := &r
	(**s).Y = 30
	//{10 30} &{10 30} 0xc042060018 0xc042060020 {10 30} &{10 30} 0xc042060018 {10 30} &{10 30} {10 30}
	fmt.Println(p, q, r, s, *q, *r, *s, **r, **s, ***s)
}
