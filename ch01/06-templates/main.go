package main

import (
	"fmt"
	"main/templates"
)

func main() {
	fmt.Println("Run in-memory text templates:")
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}

	fmt.Println("Run from-files text templates:")
	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}

	fmt.Println("Run in-memory html templates:")
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/templates	0.113s

$ go run main.go
Run in-memory text templates:

    This template demonstrates printing a "variable".


    If condition is set, we'll print this


    Next we'll iterate over an array of strings:
    0: item1
    1: item2
    2: item3


    We can also easily import other functions like strings.Split
    then immediately used the array created as a result:
    0: another_item1
    1: another_item2
    2: another_item3


    Blocks are a way to embed templates into one another
        I'm defined in a second template!




Run from-files text templates:

        Template 1! Var1!!
        Template 2! Var2!!
        Template 3! Var3!!

Run in-memory html templates:
<h1>Hello! &lt;script&gt;alert(&#39;Can you see me?&#39;)&lt;/script&gt;</h1>
JSEscaper example: \u003Cexample@example.com\u003E
HTMLEscaper example: &lt;example@example.com&gt;
URLQueryEscaper+example%3A+%3Cexample%40example.com%3E
*/
