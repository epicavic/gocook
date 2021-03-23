package channels

import "time"

// Sender sends "tick"" on ch until done is written to
func Sender(ch chan string, done chan bool) {
	t := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-done:
			ch <- "sender done." // send to receiver
			done <- true         // unblock main thread
			return
		case <-t.C:
			ch <- "tick" // send to receiver
		}
	}
}
