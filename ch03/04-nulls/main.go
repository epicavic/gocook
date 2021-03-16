package main

import (
	"fmt"

	"main/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/nulls	0.391s

$ go run main.go
jsonBlob: raw value - type: {"name": "Aaron"} - string
fulljsonBlob: raw value - type: {"name":"Aaron", "age":0} - string

jsonBlob: Regular Unmarshal, no age: {Name:Aaron Age:0}
jsonBlob: Regular Marshal, with no age: {"name":"Aaron"}
fulljsonBlob: Regular Unmarshal, with age = 0: {Name:Aaron Age:0}
fulljsonBlob: Regular Marshal, with age = 0: {"name":"Aaron"}

jsonBlob: Pointer Unmarshal, no age: {Name:Aaron Age:<nil>}
jsonBlob: Pointer Marshal, with no age: {"name":"Aaron"}
fulljsonBlob: Pointer Unmarshal, with age = 0: {Name:Aaron Age:0xc00011a230}
fulljsonBlob: Pointer Marshal, with age = 0: {"name":"Aaron","age":0}

jsonBlob: nullInt64 Unmarshal, no age: {Name:Aaron Age:<nil>}
jsonBlob: nullInt64 Marshal, with no age: {"name":"Aaron"}
fulljsonBlob: nullInt64 Unmarshal, with age = 0: {Name:Aaron Age:0xc00011a330}
fulljsonBlob: nullInt64 Marshal, with age = 0: {"name":"Aaron","age":0}
*/
