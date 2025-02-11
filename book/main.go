package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	book.Pages = pages
}

func main() {
	book1 := &Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Pages: 180}
	book2 := &Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", Pages: 281}
	book3 := &Book{Title: "Pride and Prejudice", Author: "Jane Austen", Pages: 279}

	updatePages(book1, 210)
	fmt.Println(book1)

	updatePages(book2, 250)
	fmt.Println(book2)

	updatePages(book3, 295)
	fmt.Println(book3)

	/*
		Create 3 Book Structs with the following data:

		Book 1:
		Title: "The Great Gatsby"
		Author: "F. Scott Fitzgerald"
		Pages: 180

		Book 2
		Title: "To Kill a Mockingbird"
		Author: "Harper Lee"
		Pages: 281

		Book 3
		Title: "Pride and Prejudice"
		Author: "Jane Austen"
		Pages: 279
	*/

	// your code for creating struct objects goes here

	/*
		Update the information for Books as following:

		Book 1: Updates Page Count to 210
		Book 2: Updates Page Count to 250
		Book 3: Updates Page Count to 295

	*/

	// your code for function calls to updatePages goes here

	/*
		Print all the struct objects
		fmt.Println(book)
	*/

	// your code for printing objects goes here
}
