package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(raw string) string {
	var hash = md5.Sum([]byte(raw))
	return hex.EncodeToString(hash[:])
}

func main() {
	var password string

	fmt.Println("Type your pasword: ")
	fmt.Scanln(&password)

	hash := encodeWithMD5(password)

	fmt.Println("Raw", password)
	fmt.Println("Hash", hash)
}
