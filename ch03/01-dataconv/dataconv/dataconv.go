package dataconv

import "fmt"

// ShowConv demonstrates some type conversion
func ShowConv() {
	var a = 24  // int (native. int32 or int64 depending on arch)
	var b = 2.0 // float64

	// convert the int to a float64 for this calculation
	c := float64(a) * b
	fmt.Printf("float64: %v - %T\n", c, c)

	// fmt.Sprintf is a good way to convert to strings
	s := fmt.Sprintf("%.2f", b)

	// print the value and the type
	fmt.Printf("string: %s - %T\n", s, s)
}
