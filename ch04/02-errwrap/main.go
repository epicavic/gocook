package main

import (
	"fmt"

	"main/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()

	errwrap.Unwrap()
	fmt.Println()

	errwrap.StackTrace()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/errwrap	1.048s

$ go run main.go
Regular Error -  An error occurred in WrappedError: standard error
Typed Error -  An error occurred in WrappedError: typed error
Nil - <nil>

wrapped error:  wrapped: an error occurred
a typed error occurred:  wrapped: an error occurred

an error occurred
wrapped
main/errwrap.StackTrace
	.../errwrap/unwrap.go:27
main.main
	.../errwrap/main.go:16
runtime.main
	.../src/runtime/proc.go:225
runtime.goexit
	.../src/runtime/asm_amd64.s:1371
*/
