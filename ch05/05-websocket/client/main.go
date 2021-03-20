package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

// catchSig cleans up our websocket conn if we kill the program with a ctrl-c
func catchSig(ch chan os.Signal, c *websocket.Conn) {
	// block on waiting for a signal
	<-ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
	}
	return
}

func main() {
	// connect the os signal to our channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// use the ws:// Scheme to connect to the websocket
	u := "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// dispatch our signal catcher
	go catchSig(interrupt, c)

	process(c)
}

/*
$ go run .
2021/03/20 08:35:56 connecting to ws://localhost:8000/
Enter some text: test
2021/03/20 08:36:06 received back from server: "test"
Enter some text: ^C
2021/03/20 08:38:40 failed to write message: websocket: close sent
*/
