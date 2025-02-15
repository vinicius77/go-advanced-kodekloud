package main

import (
	"errors"
	"fmt"
)

/**
In the following program, throw a custom error with the given statement, for the following cases.


a.) If divisor is less than zero:

"Division is not supported for negative numbers"

b.) If divisor is equal to zero:

"Cannot divide by zero"
*/

func divideNumber(m int, n int) (int, error) {

	if n < 0 {
		return 0, errors.New("Division is not supported for negative numbers")
	}

	if n == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return m / n, nil
}

func printResult(a int, b int) {
	res, err := divideNumber(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}
}

func main() {
	a, b := 20, 10
	printResult(a, b)

	a, b = 20, -1
	printResult(a, b)

	a, b = 20, 0
	printResult(a, b)
}
