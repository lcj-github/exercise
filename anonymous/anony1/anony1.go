/**
 * Created by Administrator on 13-12-18.
 */
package main

import (
	"fmt"
)

func main() {
	var j int = 5
	//表明此匿名函数返回值的类型是func()， 即此匿名函数返回一个函数指针
	//a:= return func(){} 等价于 ，目的引入闭包概念
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}() //这个不执行

	a() //i, j: 10, 5
	j *= 2
	a() //i, j: 10, 10
}
