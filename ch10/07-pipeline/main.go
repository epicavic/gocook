package main

import (
	"context"
	"fmt"

	"main/pipeline"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	msgcount := 10
	encoderscount := 3
	printerscount := 2
	in, out := pipeline.NewPipeline(ctx, encoderscount, printerscount)

	go func() {
		for i := 0; i < msgcount; i++ {
			in <- fmt.Sprint("Message", i)
		}
	}()

	for i := 0; i < msgcount; i++ {
		<-out
	}
}
