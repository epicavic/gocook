package main

import (
	"fmt"
	"log"
	"net/http"

	"main/controllers"
)

func main() {
	c := controllers.New(&controllers.MemStorage{})
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
