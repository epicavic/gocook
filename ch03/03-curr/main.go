package main

import (
	"fmt"

	"main/currency"
)

func main() {
	// start with our user input of fifteen dollars and 93 cents
	userInput := "15.93"
	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User input 15.93 USD converted to %d pennies\n", pennies)

	// adding 15 cents
	pennies += 15
	dollars := currency.ConvertPenniesToDollarString(pennies)
	fmt.Printf("Added 15 cents, new value is %s dollars\n", dollars)
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/currency	0.105s

$ go run main.go
User input 15.93 USD converted to 1593 pennies
Added 15 cents, new value is 16.08 dollars
*/
