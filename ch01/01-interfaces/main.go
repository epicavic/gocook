package main

import (
	"bytes"
	"fmt"

	"main/interfaces"
)

/*
The Go language provides a number of I/O interfaces
It is best practice to make use of these interfaces wherever possible
rather than passing structures or other types directly

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
type ReadSeeker interface {
	Reader
	Seeker
}
*/

func main() {
	in := bytes.NewReader([]byte("Example")) // satisfies ReadSeeker interface
	out := &bytes.Buffer{}                   // satisfies Writer interface
	fmt.Print("stdout from 'Copy' func = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("string from 'out' bytes buffer =", out.String())
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/interfaces	0.442s

$ go run .
stdout from 'Copy' func = ExampleExample
string from 'out' bytes buffer = ExampleExample
*/
