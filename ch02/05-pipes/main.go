package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount takes a type that satisfies io.Reader interface and returns a map
// with each word as a key and it's number of appearances as a value
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords) // split on whitespace

	for scanner.Scan() {
		result[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}

func main() {
	fmt.Printf("string: number_of_occurrences\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}

/*
$ go test -count=1 ./...
ok  	pipes	0.107s

$ echo "1 2 3 4 5 1 2 3" | go run main.go
string: number_of_occurrences
1: 2
2: 2
3: 2
4: 1
5: 1

~/opaq/gocook/ch02/05-pipes $ go run main.go
string: number_of_occurrences
1
2
3
4
5
1
2
3        // CTRL+D is pressed here
1: 2
2: 2
3: 2
4: 1
5: 1
*/
