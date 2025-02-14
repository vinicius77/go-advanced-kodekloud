package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)

	wg.Add(2) // two go routines

	go send(ch, &wg)
	go receive(ch, &wg)

	wg.Wait()

}

func send(ch chan int, wg *sync.WaitGroup) {
	nums := [5]int{45, 27, 56, 12, 34}

	for i := 0; i < len(nums); i++ {
		ch <- nums[i]
	}

	fmt.Println("Sent all data")

	close(ch)
	wg.Done()
}

func receive(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting For The Data")
	for value := range ch {
		fmt.Println("Received", value)
	}
	wg.Done()
}
