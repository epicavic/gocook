package main

import (
	"fmt"
	"log"
	"net/http"

	"main/proxy"
)

func main() {
	p := &proxy.Proxy{
		Client:  http.DefaultClient,
		BaseURL: "https://www.golang.org",
	}
	http.Handle("/", p)
	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
