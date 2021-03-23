package channels

import (
	"context"
	"fmt"
	"time"
)

// Receiver will print anything sent on the ch chan
// and will print tock every 200 milliseconds
// this will repeat forever until a context is
// Done, i.e. timed out or cancelled
func Receiver(ctx context.Context, ch chan string) {
	t := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printer done.")
			return
		case res := <-ch:
			fmt.Println(res)
		case <-t.C:
			fmt.Println("tock")
		}
	}
}
