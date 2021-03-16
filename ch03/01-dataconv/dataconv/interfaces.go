package dataconv

import "fmt"

// CheckType uses type switch to perform several type assertions in series
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("s.(string): It's a string!")
	case int:
		fmt.Println("s.(int): It's an int!")
	default:
		fmt.Println("default: not sure what it is...")
	}
}

// Interfaces demonstrates types assertions from empty interface to concrete type
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"

	// manual check
	if val, ok := i.(string); ok {
		fmt.Println("s.(string): val is", val)
	}

	// this one should fail
	if _, ok := i.(int); !ok {
		fmt.Println("s.(int): uh oh! glad we handled this")
	}
}
