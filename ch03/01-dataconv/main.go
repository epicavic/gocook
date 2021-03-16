package main

import "main/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/dataconv	0.312s

$ go run main.go
float64: 48 - float64
string: 2.00 - string
dec: 1234 - int64
hex: 255 - int64
bool: true - bool
s.(string): It's a string!
s.(int): It's an int!
default: not sure what it is...
s.(string): val is test
s.(int): uh oh! glad we handled this
*/
