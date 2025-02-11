package main

import (
	"fmt"
	"sync"
	"time"
)

// wait groups
func calculateSquare(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	counter := 10
	start := time.Now()
	wg.Add(counter)

	for i := 0; i < counter; i++ {
		go calculateSquare(i, &wg)
	}

	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println(elapsed)

}
