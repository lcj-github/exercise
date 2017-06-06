package main

import "fmt"
import "time"

/** 达到try finally 哪样准确的Close
jack Closed !!!
lily Closed !!!
Done !

**/
type User struct {
	username string
}

func (this *User) Close() {
	fmt.Println(this.username, "Closed !!!")
}

func main() {
	u1 := &User{"jack"}
	f(u1) // 这样的方式，u1才会不依赖main函数的执行

	// 这样的方式，u2也不会依赖main函数的执行
	u2 := &User{"lily"}
	func() {
		defer u2.Close()
		// u2 do somthing
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Done !")
}

func f(u *User) {
	defer u.Close()
}
