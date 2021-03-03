package string

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shows a number of methods for searching a string
func SearchString() {
	s := "this is a test"

	// Returns true because s contains the word 'this'
	fmt.Println(strings.Contains(s, "this"))

	// Returns true because s contains the letter 'a' would also match if it contained 'b' or 'c'
	fmt.Println(strings.ContainsAny(s, "abc"))

	// Returns true because s starts with 'this'
	fmt.Println(strings.HasPrefix(s, "this"))

	// Retursn true because s ends with 'test'
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"

	// Prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// Prints "Simple String"
	fmt.Println(strings.Title(s))

	// Prints "simple string", all trailing and leading white space is removed
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// ReadString demonstrates how to create an io.Reader interface quickly with a string
func ReadString() {
	s := "simple string\n"
	r := strings.NewReader(s)

	// prints s on Stdout
	io.Copy(os.Stdout, r)
}
