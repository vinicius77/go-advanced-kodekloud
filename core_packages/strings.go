package main

import "strings"

func main() {
	text := "Go is fun"
	search := "fun"

	strings.Contains(text, search)           // true
	strings.Count(text, search)              // 1
	strings.ReplaceAll(text, search, "nice") // Go is nice
}
