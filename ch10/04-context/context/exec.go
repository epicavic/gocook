package context

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Exec sets two random timers and prints a different context value
// for whichever fires first
func Exec() {
	ctx := context.Background() // a base context
	ctx = SetContextValues(ctx) // set timeout/deadline values

	rand.Seed(time.Now().UnixNano())
	timeoutCtx, cancel := context.WithTimeout(ctx, (time.Duration(rand.Intn(2)) * time.Millisecond))
	defer cancel()

	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Duration(rand.Intn(2))*time.Millisecond))
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			fmt.Println(GetValue(ctx, timeoutKey))
			return
		case <-deadlineCtx.Done():
			fmt.Println(GetValue(ctx, deadlineKey))
			return
		}
	}
}
