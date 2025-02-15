package main

import (
	"errors"
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("Only odd numbers are allowed %d", i)
	}
	return nil
}

func main() {
	error_ := errors.New("Custom Error")
	fmt.Println(error_)

	err := process(2)
	fmt.Println(err)

	err = process(3)
	fmt.Println(err)
}
