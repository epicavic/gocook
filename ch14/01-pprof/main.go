package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/guess", GuessHandler)
	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}

/*
func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
	http.HandleFunc("/debug/pprof/trace", Trace)
}

$ go tool pprof http://localhost:8080/debug/pprof/profile
$ go tool pprof http://localhost:8080/debug/pprof/heap
*/
