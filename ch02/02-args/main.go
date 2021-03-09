package main

import (
	"fmt"
	"os"
	"strings"

	"main/args"
)

func main() {
	c := args.MenuConf{}
	menu := c.SetupMenu()

	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	// we use arguments to switch between commands
	// flags are also arguments
	if len(os.Args) > 1 {
		// we don't care about case
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing params %s, error: %v", os.Args[3:], err)
					return
				}

			}
			c.Greet(os.Args[2])

		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}
	} else {
		menu.Usage()
		return
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/args	0.272s

$ go run main.go
Usage:
main [command]

Commands:
    greet
    version

$ go run main.go version
Version: 1.0.0

$ go run main.go greet
Usage:
main greet name [flag]

Positional Arguments:
    name
        the name to greet

Flags:
  -goodbye
    	Say goodbye instead of hello

$ go run main.go greet vic
Hello vic!

$ go run main.go greet vic -goodbye
Goodbye vic!
*/
