package main

import "main/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}

/*
$ go test -count=1 ./...
?   	main	[no test files]
ok  	main/tempfiles	0.807s

$ go run main.go
Temp dir: /var/folders/bz/xvk1v1vx7vb99ts2g3jv72gh0000gn/T/tmp833797689
Temp file: /var/folders/bz/xvk1v1vx7vb99ts2g3jv72gh0000gn/T/tmp833797689/tmp588692036
*/
