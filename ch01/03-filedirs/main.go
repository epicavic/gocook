package main

import "main/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.Capitalizer(); err != nil {
		panic(err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/filedirs	0.087s

$ go run main.go
hello
[example_dir filedirs go.mod main.go]
*/
