package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

//{1 2} &{1 2} {1 0} {0 0}
func main() {
	fmt.Println(p, q, r, s)
	fmt.Printf("%T %T %T %T", p, q, r, s) // main.Vertex *main.Vertex main.Vertex main.Vertex
}
