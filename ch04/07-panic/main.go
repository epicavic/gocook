package main

import (
	"fmt"

	"main/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}

/*
$ go run main.go
before panic
panic occurred: runtime error: integer divide by zero
after panic
*/
