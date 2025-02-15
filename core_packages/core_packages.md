##### Core Packages (Built-in)

- `string`

  - `Contains(string, string) boolean`
  - `ReplaceAll(string, string, string) string`
  - `Count(string, string) int`

  ```go
  package main
  import "strings"

  func main(){
    text:= "Go is fun"
    search: = "fun"

    strings.Contains(text, fun) // true
    strings.Count(text, search) // 1
    strings.ReplaceAll(text, search, "nice") // Go is nice
  }
  ```

- Input / Output (`io`)

  - `Reader`

  ```go
  import (
    "strings"
  )

  func main(){
    r := strings.NewReader("Learning is fun")
    buf := make([]byte, 4)

    n, err := r.Reader(buf)
    fmt.Println(n, buf[:n], err, string(buf[:n])) // 4 [76, 101, 97, 14] nil Lear

    for {
      n, err := r.Read(buf)
      fmt.Println(string(buf[:n], err)) // "lear" nil, "ning" nil, " is " nil, "fun" nil, EOF
      if err != nil {
        fmt.Println("Breaking out")
        break
      }
    }
  }
  ```

  - `Writer`

  ```go
  import (
    "io"
    "strings"
    "log"
    "os"
  )

  func main(){
    r := strings.NewReader("some io.Reader stream to be read \n")

    // /dev/stdout
    if _, err := io.Copy(os.Stdout, r); err != nil {
      log.Fatal(err)
    }
  }
  ```

- File Handling

  ```go
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
  ```

- Errors `errors`

  - `New`
  - `Errorf`

  ```go
  package main

  import (
    "errors"
    "fmt"
  )

  func process(i int) error {
    if i%2 == 0 {
      return fmt.Errorf("Only odd numbers are allowed %d", i)
    }
    return nil
  }

  func main() {
    error_ := errors.New("Custom Error")
    fmt.Println(error_)

    err := process(2)
    fmt.Println(err)

    err = process(3)
    fmt.Println(err)
  }
  ```

- Logging `log`

  ```go
    package main

    import (
      "fmt"
      "log"
      "os"
    )

    // Glog or Logrus (famous)
    // github.com/sirupsen/logrus
    // - Trace, Debug, Info, Warn, Error, Fatal, Panic i.e log.Trace(), log.Debug() etc
    func main() {
      logFile := "/media/kako77sub/DA182A9E182A79A11/Users/uchiha-kako/projects/projects/go-advanced-kodekloud/log.txt"
      f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      defer f.Close()
      // logrus.SetOutput(f)
      log.SetOutput(f)
      log.Println("Some logs happening here")
    }
  ```

- Sort

  ```go
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
  ```

- Hashes and Cryptography

```go
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
```

- Testing
