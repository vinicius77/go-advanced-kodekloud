package main

/**
We are working on a project that involves
processing a large number of text documents.
We need to write a function that counts the
number of times a specific word appears in
a given string.

Let's write a function WordCount that takes in a string s and a string word and returns the number of times the word appears in s. The function should be case-sensitive, so hello and Hello are considered to be different words.
*/

import (
	"fmt"
	"strings"
)

func WordCount(s string, word string) int {
	return strings.Count(s, word)
}

func main() {
	count := WordCount("hello, Hello how have you been in helloworld", "hello")
	fmt.Println(count)
}
