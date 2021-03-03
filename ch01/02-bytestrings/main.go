package main

import (
	mb "main/buffer"
	ms "main/string"
)

func main() {
	err := mb.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	ms.SearchString()
	ms.ModifyString()
	ms.ReadString()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/buffer	0.085s
ok  	main/string	0.140s

$ go run main.go
it's easy to encode unicode into a byte array ❤️
it's easy to encode unicode into a byte array ❤️
it'seasytoencodeunicodeintoabytearray❤️true
true
true
true
[simple string]
Simple String
simple string
simple string
*/
