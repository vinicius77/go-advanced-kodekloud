// Package writer provides a simple example of reading from an io.Reader and writing to standard output (os.Stdout).
package writer

import (
	"io"
	"log"
	"os"
	"strings"
)

// main reads a string from an io.Reader stream and writes it to standard output (os.Stdout).
// If an error occurs during writing, it logs the error and terminates the program.
// TLDR: Writes the string into the terminal console, similar to fmt.Println("some io.Reader stream to be read")
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	// /dev/stdout
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
