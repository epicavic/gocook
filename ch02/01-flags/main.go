package main

import (
	"flag"
	"fmt"

	"main/flags"
)

func main() {
	// initialize our setup
	c := flags.Config{}
	c.Setup()

	// parse flags/values and arguments
	// flag package requires all flags to appear before positional arguments
	// (otherwise the flags will be interpreted as positional arguments)
	flag.Parse()

	fmt.Println(c.GetMessage())
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/flags	0.391s

$ go run main.go -isawesome=false -s Luke -howawesome 8 -c "1,2,3,4,5,6,7,8,9"
Luke is NOT awesome with a certainty of 8 out of 10. Let me count the ways 1 ... 2 ... 3 ... 4 ... 5 ... 6 ... 7 ... 8 ... 9
*/
