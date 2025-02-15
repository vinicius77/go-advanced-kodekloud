package main

import (
	"fmt"
	"os"
)

/**
A text file is given called temp.txt.
Write a code in main.go file to open
and read from the file in such a way,
to produce the following output:

Lorem Ipsum is simply dummy text of the printing a

nd typesetting industry. Lorem Ipsum has been the

industry's standard dummy text ever since the 1500

s, when an unknown printer took a galley of type a

nd scrambled it to make a type specimen book.

Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.
*/

func main() {
	f, _ := os.Open("/media/kako77sub/DA182A9E182A79A11/Users/uchiha-kako/projects/projects/go-advanced-kodekloud/exercises/temp.txt")
	b := make([]byte, 50)

	for {
		n, err := f.Read(b)
		fmt.Println(string(b[0:n]))

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
