package main

import (
	"fmt"
	"log"
	"net/http"

	"main/negotiate"
)

func main() {
	http.HandleFunc("/", negotiate.Negotiate)
	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
