package math

import (
	"fmt"
	"math"
)

// Examples demonstrates some of the functions in the math package
func Examples() {
	// sqrt. i is an int, so convert
	i := 25
	result := math.Sqrt(float64(i))
	fmt.Println("sqrt(25): ", result)

	// ceil rounds up
	result = math.Ceil(9.5)
	fmt.Println("ceil(9.5): ", result)

	// floor rounds down
	result = math.Floor(9.5)
	fmt.Println("floor(9.5): ", result)

	// math also stores some consts:
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}
