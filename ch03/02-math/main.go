package main

import (
	"fmt"

	"main/math"
)

func main() {
	math.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/math	1.015s

$ go run main.go
sqrt(25):  5
ceil(9.5):  10
floor(9.5):  9
Pi: 3.141592653589793 E: 2.718281828459045
1 1 2 3 5 8 13 21 34 55
*/
