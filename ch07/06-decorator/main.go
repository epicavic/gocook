package main

import "main/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
