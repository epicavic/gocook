package main

import (
	"fmt"

	"main/color"
)

func main() {
	r := color.ColorText{
		TextColor: color.Red,
		Text:      "I'm red!",
	}

	fmt.Println(r.String())

	r.TextColor = color.Green
	r.Text = "Now I'm green!"

	fmt.Println(r.String())

	r.TextColor = color.ColorNone
	r.Text = "Back to normal..."

	fmt.Println(r.String())
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/color	0.764s

$ go run main.go
I'm red!
Now I'm green!
Back to normal...
*/
