package reader

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Learning is fun")
	buf := make([]byte, 4)

	n, err := r.Read(buf)
	fmt.Println(n, buf[:n], err, string(buf[:n])) // 4 [76, 101, 97, 14] nil Lear

	for {
		n, err := r.Read(buf)
		fmt.Println(string(buf[:n]), err) // "lear" nil, "ning" nil, " is " nil, "fun" nil, EOF

		if err != nil {
			fmt.Println("Breaking out")
			break
		}
	}
}
