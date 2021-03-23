package main

import (
	"context"
	"time"

	"main/channels"
)

func main() {
	ch := make(chan string) // message channel used for sending string messages
	done := make(chan bool) // signaling channel used for synchronization

	ctx := context.Background()            // init background context
	ctx, cancel := context.WithCancel(ctx) // init context cancel function

	go channels.Sender(ch, done)  // dispatch sender gorputine
	go channels.Receiver(ctx, ch) // dispatch receiver goroutine

	time.Sleep(2 * time.Second) // send and receive for 2 sec
	done <- true                // stop sender
	<-done                      // block until sender has finished
	cancel()                    // stop receiver
	time.Sleep(1 * time.Millisecond)
}

/*
$ go run main.go
tock
tick
tock
tick
tock
tick
sender done.
printer done.
*/
