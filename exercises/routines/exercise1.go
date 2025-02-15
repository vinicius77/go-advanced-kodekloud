package main

import (
	"fmt"
	"sync"
)

/**
Complete the code below to create two goroutines
that will run concurrently using a WaitGroup.
The first goroutine should print the numbers 1 to 5,
and the second goroutine should print the letters a to e.

The main function should wait for both goroutines to
finish before exiting. */

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	// Create first goroutine here
	go printNumbers(&wg)
	// Create second goroutine here
	go printLetters(&wg)

	// Wait for both goroutines to finish here
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	numbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	letters := [5]string{"a", "b", "c", "d", "e"}

	for _, letter := range letters {
		fmt.Println(letter)
	}

}
