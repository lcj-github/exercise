package main

import "fmt"

func main() {

	mesnoBuffer := make(chan string)
	go func() { mesnoBuffer <- "ping" }() //start a goroutine for an anonymous function call.
	msg := <-mesnoBuffer
	fmt.Println(mesnoBuffer) //0xc042030180
	fmt.Println(msg)         // ping

	mesHaveBuffer := make(chan string, 2)
	mesHaveBuffer <- "buffered"
	mesHaveBuffer <- "channel"
	fmt.Println(<-mesHaveBuffer) //buffered
	fmt.Println(<-mesHaveBuffer) //channel
}
