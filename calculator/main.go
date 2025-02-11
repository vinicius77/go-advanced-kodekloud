package main

import "fmt"

/*
In this question, we are building a simple calculator
that can perform addition, subtraction, multiplication, and division operations.

Complete the code in the function calculate to return a slice consisting of 4
elements
[ sum of a and b, difference of a and b, product of a and b, quotient on dividing a by b]

A Go file is located at /root/code/calculator/ directory.
*/

func calculate(num1 int, num2 int) []float64 {
	sum := float64(num1 + num2)
	sub := float64(num1 - num2)
	prod := float64(num1 * num2)
	div := float64(num1 / num2)

	res := []float64{sum, sub, prod, div}

	return res

}

func main() {
	fmt.Println(calculate(20, 10))

	fmt.Println(calculate(700, 70))
}
