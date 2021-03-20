package main

import (
	"fmt"
	"log"
	"net/rpc"

	"main/tweak"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	args := tweak.Args{
		String:  "this string should be uppercase and reversed",
		ToUpper: true,
		Reverse: true,
	}
	var result string
	err = client.Call("StringTweaker.Tweak", args, &result)
	if err != nil {
		log.Fatal("client call with error:", err)
	}
	fmt.Printf("the request  is: %+v\n", args)
	fmt.Printf("the responce is: %s\n", result)
}

/*
$ go run client/main.go
the request  is: {String:this string should be uppercase and reversed ToUpper:true Reverse:true}
the responce is: DESREVER DNA ESACREPPU EB DLUOHS GNIRTS SIHT
*/
