package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

/**
We have given a text file called input.txt that
contains a list of integers, one per line. Our
task is to read the given file, sum the integers,
and write the sum to a new file called output.txt
within a same directory.
*/

func main() {
	// your code goes here
	// Open the input file
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	path := filepath.Join(cwd, "exercises/xerox", "input.txt")

	inputFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Create a scanner to read the input file
	inputScanner := bufio.NewScanner(inputFile)

	// Initialize the sum to 0
	sum := 0

	// Iterate over the lines of the input file
	for inputScanner.Scan() {
		// Parse the current line as an integer
		n, err := strconv.Atoi(inputScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		// Add the integer to the sum
		sum += n
	}

	// Check for errors while reading the input file
	if err := inputScanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Open the output file
	path = filepath.Join(cwd, "exercises/xerox", "output.txt")
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Write the sum to the output file as a string
	_, err = outputFile.WriteString(strconv.Itoa(sum))
	if err != nil {
		fmt.Println(err)
		return
	}

}
