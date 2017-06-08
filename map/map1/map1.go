package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

//make 函数构造 map，貌似不能通过 new ；
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	m["abc"] = Vertex{
		1.2, 3.4,
	}
	fmt.Println(m["abc"])
}
