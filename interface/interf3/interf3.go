package main

import (
	"fmt"
	"os"
)

/**
1). 接口继承；
2). os 包；os.Stdout ； fmt.Fprintf ；标准输出流、流格式化输出；
os.Stdout 实现了 Writer 接口，可以使用 fmt.Fprintf 向 os.Stdout 写内容；
**/
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n") //hello, writer
}
