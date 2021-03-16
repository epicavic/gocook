package main

import (
	"main/encoding"
)

func main() {
	if err := encoding.Base64Example(); err != nil {
		panic(err)
	}

	if err := encoding.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := encoding.GobExample(); err != nil {
		panic(err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/encoding	0.249s

$ go run main.go
Source string:  encoding some data!
With URLEncoding.EncodeToString:  ZW5jb2Rpbmcgc29tZSBkYXRhIQ==
With URLEncoding.DecodeToString:  encoding some data!
With NewEncoder and StdEncoding:  ZW5jb2Rpbmcgc29tZSBkYXRhIQ==
With NewDecoder and StdEncoding:  encoding some data!
Gob Encoded value length - type: 57 - *gob.Encoder
Gob Decoded value - type:  {10 15 wrench} - *gob.Decoder
*/
