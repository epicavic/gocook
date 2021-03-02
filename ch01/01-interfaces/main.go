package main

import (
	"bytes"
	"fmt"

	"main/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("Example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout from 'Copy' func = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("string from 'out' bytes buffer =", out.String())
}
