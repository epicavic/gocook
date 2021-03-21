package main

import "main/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
