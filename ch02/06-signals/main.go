package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig sets up a listener catching signals
func CatchSig(ch chan os.Signal, done chan bool) {
	// block on waiting for a signal
	sig := <-ch
	// print it when it's received
	fmt.Println("sig received:", sig)

	// set up handlers here
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM now!")
	default:
		fmt.Println("unexpected signal received")
	}

	// terminate
	done <- true
}

func main() {
	// initialize our channels
	signals := make(chan os.Signal)
	done := make(chan bool)

	// hook them up to the signals lib
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// if a signal is caught by this go routine it will write to done
	go CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	// the program blocks until someone writes to done
	<-done
	fmt.Println("Done!")
}

/*
$ go test -count=1 ./...
ok  	signals	0.381s

$ (go run main.go &); sleep 1 && kill -SIGINT $(ps -ef | grep [e]xe/main | awk '{print $2}')
Press ctrl-c to terminate...
sig received: interrupt
handling a SIGINT now!
Done!

$ (go run main.go &); sleep 1 && kill -SIGTERM $(ps -ef | grep [e]xe/main | awk '{print $2}')
Press ctrl-c to terminate...
sig received: terminated
handling a SIGTERM now!
Done!
*/
