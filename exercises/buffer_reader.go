package main

/**
Write a code in main.go file to read the string in
a buffer of size 5 at a time and print the output
in each iteration.

A Go file is located at /root/code/string directory. */

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Let us catch up over a cup of coffee")
	// your code goes here
	buf := make([]byte, 5)

	for {
		n, err := reader.Read(buf)
		fmt.Println(buf[:n], err)

		if err != nil {
			break
		}
	}

}
