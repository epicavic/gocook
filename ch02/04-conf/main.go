package main

import "main/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/confformat	0.117s

$ go run main.go
TOML Marshal:
name = "Name1"
age = 20

YAML Marshal:
name: Name3
age: 40

JSON Marshal:
{"name":"Name2","age":30}

TOML Unmarshal = {Example1 99}
Yaml Unmarshal = {Example3 97}
JSON Unmarshal = {Example2 98}

We can unmarshal into a map instead of a struct: map[key1:value1]
We can also use decoders/encoders to work with streams: map[key1:value1 key2:value2]
*/
