package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go sell(ch, &wg)
	go buy(ch, &wg)

	wg.Wait()

}

// unbuffered channel, does not store values

// sends data
// channel,  slice and maps are sent as reference implicity (no need for *)
func sell(ch chan string, wg *sync.WaitGroup) {
	ch <- "Channel string data" // blocks the 'sell' method until another go routine receives it
	fmt.Println("Data sent")
	wg.Done()
}

// receive data
func buy(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	value := <-ch
	fmt.Println("Receiving data -", value)
	wg.Done()
}
