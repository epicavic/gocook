package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"main/envs"
)

// Config struct will hold the config we capture from a json file and env vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	// create temp file to hold a json
	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	// create a json with secret and write to temp file
	secrets := `{
        "secret": "super secret"
    }`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// set environment variables as needed
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	// c will hold config values
	c := Config{}
	if err = envs.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("json file contains =", secrets)

	// We can also read them
	fmt.Println("env var EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("env var EXAMPLE_ISSAFE  =", os.Getenv("EXAMPLE_ISSAFE"))

	// The final config is a mix of json and environment variables
	fmt.Printf("\nFinal Config:\n%#v\n", c)

}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/envs	0.992s

$ go run main.go
json file contains = {
        "secret": "super secret"
    }
env var EXAMPLE_VERSION = 1.0.0
env var EXAMPLE_ISSAFE  = false

Final Config:
main.Config{Version:"1.0.0", IsSafe:false, Secret:"super secret"}
*/
