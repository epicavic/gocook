package main

import (
	"fmt"

	"main/structured"
)

func main() {
	fmt.Println("Logrus in text:")
	structured.Logrus("text")
	fmt.Println()

	fmt.Println("Logrus in json:")
	structured.Logrus("json")
	fmt.Println()

	fmt.Println("Apex:")
	structured.Apex()
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/structured	1.003s

$ go run main.go
Logrus in text:
WARN[0000] warning!                                      complex_struct="{Something happened Just now}" id=123 success=true
ERRO[0000] error!                                        complex_struct="{Something happened Just now}" id=123 success=true

Logrus in json:
{"complex_struct":{"Event":"Something happened","When":"Just now"},"id":"123","level":"warning","msg":"warning!","success":true,"time":"2021-03-18T16:15:27+02:00"}
{"complex_struct":{"Event":"Something happened","When":"Just now"},"id":"123","level":"error","msg":"error!","success":true,"time":"2021-03-18T16:15:27+02:00"}

Apex:
  INFO[0000] ThrowError                id=123
 ERROR[0000] ThrowError                duration=0 error=a crazy failure id=123
 ERROR[0000] an error occurred         error=a crazy failure
*/
