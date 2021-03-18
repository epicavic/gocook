package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue is a way to make a package level error to check against.
// I.e. if err == ErrorValue
var ErrorValue = errors.New("this is an error value")

// TypedError is a way to make an error type
// I.e. err.(type) == TypedError
type TypedError struct {
	error // 'Error() string' method promoted from error interface type
}

//BasicErrors demonstrates some ways to create errors
func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Printf("errors.New: %v - %T\n", err, err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Printf("fmt.Errorf: %v - %T\n", err, err)

	err = ErrorValue
	fmt.Printf("value error: %v - %T\n", err, err)

	err = TypedError{errors.New("typed error")}
	fmt.Printf("typed error: %v - %T\n", err, err)

	switch err.(type) {
	case TypedError:
		fmt.Println("error matched: TypedError type")
	}
}
