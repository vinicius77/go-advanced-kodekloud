package main

import (
	"fmt"
	"sync"
)

// block the go routine only if the buffer is full
func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)

	wg.Add(2)
	go sell(ch, &wg)

	wg.Wait()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for the data")
	fmt.Println("Received data:", <-ch)
	wg.Done()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	ch <- 3
	buy(ch, wg)

	fmt.Println("Sent all data to the channel:", len(ch))
	wg.Done()

}
