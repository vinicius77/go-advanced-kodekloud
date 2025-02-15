package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// go doc filepath Join
	path := filepath.Join("dir1", "dir2", "text.txt")

	fmt.Println(path)                          // dir1/dir2/text.txt
	fmt.Println(filepath.Dir(path))            // dir1/dir2
	fmt.Println(filepath.Base(path))           // text.txt
	fmt.Println(filepath.IsAbs(path))          // false => is Absolute path?
	fmt.Println(filepath.IsAbs("/dev/stdout")) // true
	fmt.Println(filepath.Ext(path))            // .txt

	// ==== File Reader
	validFilePath := "/media/kako77sub/DA182A9E182A79A11/Users/uchiha-kako/projects/projects/go-advanced-kodekloud/core_packages/writer.go"
	invalidFilePath := "../write.go"

	fileInfo, err := os.Stat(invalidFilePath)
	fmt.Println(fileInfo, err) // <nil> stat ../write.go: no such file or directory

	fileInfo, err = os.Stat(validFilePath)
	// &{writer.go 631 511 {354488700 63875233914 0x537a00} {2052 514905 1 33279 1000 1000 0 0 631 4096 2 {1739637114 407481400} {1739637114 354488700} {1739637114 354488700} [0 0 0]}} <nil>
	fmt.Println(fileInfo, err)
	fmt.Println(fileInfo.Name())  // writer.go
	fmt.Println(fileInfo.Size())  // 631
	fmt.Println(fileInfo.Mode())  // -rwxrwxrwx
	fmt.Println(fileInfo.IsDir()) // false

	data, err := os.ReadFile(validFilePath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)         // ascii characters
	fmt.Println(string(data)) // file content

	// read big files in batches
	file, err := os.Open(validFilePath)

	if err != nil {
		fmt.Println(err)
	}

	b := make([]byte, 40)

	for {
		n, err := file.Read(b)

		if err != nil {
			fmt.Println("Error:", err) // Error: EOF
			break
		}

		fmt.Println(b[:n], string(b[:n]))
	}

	// Write / Append to a file
	canvasFilePath := "/media/kako77sub/DA182A9E182A79A11/Users/uchiha-kako/projects/projects/go-advanced-kodekloud/canvas.txt"
	f, err := os.OpenFile(canvasFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println("Error:", err) // Error: EOF
		return
	}

	defer f.Close()

	currentTime := time.Now().Unix()
	formattedString := fmt.Sprintf("File written at: %v\n", currentTime)
	_, err = f.WriteString(formattedString)

	fmt.Println("File opened successfully")

}
