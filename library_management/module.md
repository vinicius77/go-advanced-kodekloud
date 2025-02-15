##### Go Module

- To create a module (path must be globally unique)

```go
go mod init github.com/vinicius77/go-advanced-kodekloud
```

```
go: creating new go.mod: module github.com/vinicius77/go-advanced-kodekloud
go: to add module requirements and sums:
	go mod tidy
```

- Check example `go.mod` file on project's root

```go
module github.com/vinicius77/go-advanced-kodekloud

go 1.18
```

When adding external packages, run `go mod tidy`, the `go.mod` file is going to be updated with the new dependencies

```bash
$: go mod tidy

go: finding module for package github.com/pborman/uuid
go: downloading github.com/pborman/uuid v1.2.1
go: found github.com/pborman/uuid in github.com/pborman/uuid v1.2.1
go: downloading github.com/google/uuid v1.0.0
```

```go
module github.com/vinicius77/go-advanced-kodekloud

go 1.18

require github.com/pborman/uuid v1.2.1

require github.com/google/uuid v1.0.0 // indirect
```

If trying to import local modules, for example, that were not yet published in any repository: (Note the `=`), run the command:

```bash
go mod edit -replace <module_path>=<local_module_patch>
```

The `go.mod` file will look like similar to:

```go
module github.com/vinicius77/go-advanced-kodekloud

go 1.18

// .. rest of the dependencies here

replace github.com/vinicius77/go-advanced-kodekloud/encode => ../encode
```

###### Go Commands (most commom)

- `go mod init`
- `go mod tidy`
- `go run <file>`
- `go build` (complies the packages into an executable)
- `go install` (compiles and moves the executable to the GO_PATH/bin )
- `go get` (add, upgrade packages)

#### Compiling and placing executable into /bin

```bash
go env GOPATH
/home/kako77sub/go

export PATH=$PATH:/home/kako77sub/go
```

- In the project path

```go
go install
```

#### Publishing a module

- Use versioning standards (semver) `git tag vx.x.x`
- Pull the project to a repository (Gitlab, Github, Bitbucket etc)
- Remove `replace` (local module) and run `go get github.com/vinicius77/go-advanced-kodekloud/encode@vx.x.x`

Documentation (`godoc` tool)

```go
// go doc PACKAGE_NAME
go doc encode

// go doc PACKAGE_NAME.IDENTIFIER_NAME
go doc encode.decrypt
```

#### Avoiding Package name collision

```go
import (
	"math/rand"
	crand "crypto/rand"
)


func main(){
	crand.Read(/* ... */)
	rand.New(/* ... */)
}
```

#### Naming packages

- Good

```go
format.Names
extract.Names
```

- Bad

```go
util.FormatNames
util.ExtractNames
```

```go
format.FormatNames
format.ExtractNames
```
