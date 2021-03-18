package main

import (
	"fmt"

	"main/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()

	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/log	0.363s

$ go run main.go
basic logging and modification of logger:
new logger: 2021/03/18 log.go:18: test
set prefix: 2021/03/18 log.go:22: you can also add args(true) and use Fatalln to log and crash

logging 'handled' errors:
2021/03/18 14:32:14 an error occurred: in passthrougherror: error occurred
*/
