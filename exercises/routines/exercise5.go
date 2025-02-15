package main

import (
	"fmt"
	"sync"
)

/**
In this question, we are working on a
project that involves processing a
large amount of data in parallel using
goroutines. We want to use channels to
communicate between the goroutines and
ensure that all goroutines have finished
processing before the program exits.

Complete the code, so that the program runs
successfully without any errors.
*/

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for j := range jobs {
			fmt.Println("received job", j)
			result := j * 2
			results <- result
		}
		// your code goes here
		wg.Done()
		close(results)
	}()

	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		// your code goes here
		wg.Done()
		close(jobs)
	}()

	go func() {
		for r := range results {
			fmt.Println("received result", r)
		}
		// your code goes here
		wg.Done()
	}()

	wg.Wait()
}
