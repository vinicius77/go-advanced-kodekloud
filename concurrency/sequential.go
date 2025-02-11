package main

import (
	"fmt"
	"time"
)

// sequential program
func calculateSquare(num int) {
	time.Sleep(1 * time.Second)
	fmt.Println(num * num)
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		calculateSquare(i)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
