package main

import "main/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
