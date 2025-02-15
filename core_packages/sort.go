package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []int{5, 243, 2, 89, 54, 92}
	strings := []string{"car", "bat", "zebra", "bee", "dog"}

	sort.Ints(vars)
	sort.Strings(strings)

	fmt.Println(vars, strings)
}
