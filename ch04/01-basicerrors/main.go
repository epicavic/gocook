package main

import (
	"fmt"

	"main/basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	fmt.Println()

	if err := basicerrors.SomeFunc(); err != nil {
		fmt.Println("custom error:", err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/basicerrors	0.218s

$ go run main.go
errors.New: this is a quick and easy way to create an error - *errors.errorString
fmt.Errorf: an error occurred: something - *errors.errorString
value error: this is an error value - *errors.errorString
typed error: typed error - basicerrors.TypedError
error matched: TypedError type

custom error: there was an error; this was the result
*/
