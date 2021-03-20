package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on localhost:8000")
	// we mount our single handler on port localhost:8000 to handle all requests
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(wsHandler)))
}

/*
$ go run .
Listening on localhost:8000
2021/03/20 08:36:06 received from client: "test"
2021/03/20 08:38:37 failed to read message:  websocket: close 1000 (normal)
^Csignal: interrupt
*/
