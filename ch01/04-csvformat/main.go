package main

import (
	"fmt"

	"main/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}
	fmt.Println()

	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Buffer =\n%v", buffer.String())
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/csvformat	0.192s

$ go run main.go
[]csvformat.Movie{csvformat.Movie{Title:"Guardians of the Galaxy Vol. 2", Director:"James Gunn", Year:2017}, csvformat.Movie{Title:"Star Wars: Episode VIII", Director:"Rian Johnson", Year:2017}}

Author,Title
F Scott Fitzgerald,The Great Gatsby
J D Salinger,The Catcher in the Rye

Buffer =
Author,Title
F Scott Fitzgerald,The Great Gatsby
J D Salinger,The Catcher in the Rye
*/
