package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"main/tweak"
)

func main() {
	s := new(tweak.StringTweaker)
	if err := rpc.Register(s); err != nil {
		log.Fatal("failed to register:", err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Println("listening on localhost:1234")
	log.Panic(http.Serve(l, nil))
}

/*
$ go run server/main.go
listening on localhost:1234
*/
