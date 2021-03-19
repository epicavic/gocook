package main

import "main/context"

func main() {
	context.Initialize()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/context	0.369s

$ go run main.go
  INFO[0000] starting                  id=123
  INFO[0000] after gatherName          id=123 name=Go Cookbook
  INFO[0000] after gatherLocation      city=Seattle id=123 name=Go Cookbook state=WA
*/
