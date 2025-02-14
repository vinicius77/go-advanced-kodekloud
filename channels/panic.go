package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	ch <- 10
	ch <- 15

	val, ok := <-ch
	fmt.Println(val, ok) // 10, true
	close(ch)

	// ch <- 20 // panic: send on closed channel
	close(ch) // panic: close of closed channel

}
