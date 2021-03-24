package main

import (
	"fmt"

	"main/pool"
)

func main() {
	c := 10                             // worker pool capacity
	cancel, in, out := pool.Dispatch(c) // create worker pool and cancellation function
	defer cancel()

	for i := 0; i < c; i++ {
		in <- pool.WorkRequest{Op: pool.Hash, Text: []byte(fmt.Sprintf("messages %d", i))}
	}

	for i := 0; i < c; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- pool.WorkRequest{Op: pool.Compare, Text: res.Wr.Text, Compare: res.Result}
	}

	for i := 0; i < c; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string: \"%s\"; matched: %v\n", string(res.Wr.Text), res.Matched)
	}
}
