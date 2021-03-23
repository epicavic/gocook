package main

import (
	"fmt"
	"log"

	"main/waitgroup"
)

func main() {
	sites := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://www.google.com/search?q=golang",
	}

	resps, err := waitgroup.Crawl(sites)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Resps received:", resps)
}

/*
$ go run main.go
2021/03/23 14:18:17 starting crawling
2021/03/23 14:18:17 getting https://www.google.com/search?q=golang
2021/03/23 14:18:17 getting https://godoc.org
2021/03/23 14:18:17 getting https://golang.org
2021/03/23 14:18:18 completed getting https://www.google.com/search?q=golang in 408.502445ms
2021/03/23 14:18:18 completed getting https://golang.org in 477.506022ms
2021/03/23 14:18:18 completed getting https://godoc.org in 860.281877ms
2021/03/23 14:18:18 completed crawling in 860.72638ms
Resps received: [200 200 200]
*/
