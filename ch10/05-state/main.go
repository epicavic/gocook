package main

import (
	"context"
	"fmt"

	"main/state"
)

func main() {
	reqs := []state.WorkRequest{
		{Op: state.Add, V1: 3, V2: 4},
		{Op: state.Subtract, V1: 5, V2: 2},
		{Op: state.Multiply, V1: 9, V2: 9},
		{Op: state.Divide, V1: 8, V2: 2},
		{Op: state.Divide, V1: 8, V2: 0},
	}
	in := make(chan state.WorkRequest, len(reqs))
	out := make(chan state.WorkResponse, len(reqs))
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processor(ctx, in, out)

	for _, req := range reqs {
		in <- req
	}

	for range reqs {
		resp := <-out
		fmt.Printf("Request: %v; Result: %v\n", resp.Wr, resp.Result)
	}
}
