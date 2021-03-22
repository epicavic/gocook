package decorator

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Exec creates a client, calls google.com then prints the response
func Exec() error {
	c := Setup()

	resp, err := c.Get("https://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("Response code:", resp.StatusCode)
	return nil
}

// Setup initializes our ClientInterface
func Setup() *http.Client {
	c := http.Client{}

	t := Decorate(&http.Transport{},
		Logger(log.New(os.Stdout, "", 0)),
		BasicAuth("username", "password"),
	)
	c.Transport = t
	return &c
}
