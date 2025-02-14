package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	ch <- 10
	ch <- 15

	val, ok := <-ch
	fmt.Println(val, ok) // 10, true
	close(ch)

	val, ok = <-ch
	fmt.Println(val, ok) // 15, true

	val, ok = <-ch
	fmt.Println(val, ok) // 0, false
}
