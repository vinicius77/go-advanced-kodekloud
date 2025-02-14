package main

import (
	"fmt"
	"sync"
)

// It seems similar to race for Redux Sagas
func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(1) // one go routine is needed without the default case

	go channel1(ch1, &wg)
	go channel2(ch2, &wg)

	select {
	case value1 := <-ch1:
		fmt.Println(value1)
	case value2 := <-ch2:
		fmt.Println(value2)
		//default: // non blocking main function
		//	wg.Done()
		//	fmt.Println("Executed the default")
		//	break
	}

	wg.Wait()

}

func channel1(ch1 chan string, wg *sync.WaitGroup) {
	ch1 <- "Channel One Data"
	close(ch1)
	wg.Done()
}

func channel2(ch2 chan string, wg *sync.WaitGroup) {
	ch2 <- "Channel Two Data"
	close(ch2)
	wg.Done()
}
