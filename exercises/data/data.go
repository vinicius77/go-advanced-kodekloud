package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

/**
We are working on a project that involves reading data
from a text file and processing it. We need to write a
function that reads the contents of a text file and
returns a slice of strings, where each string represents
a line in the file
*/

func ReadLines(filename string) ([]string, error) {
	// your code goes here
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	path := filepath.Join(cwd, "exercises/data", "data.txt")
	lines, err := ReadLines(path)

	if err != nil {
		fmt.Println(err)
	} else {

		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
