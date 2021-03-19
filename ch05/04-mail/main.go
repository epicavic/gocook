package main

import (
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

// an example email message
const msg string = `Date: Thu, 18 Mar 2021 08:00:00 -0700
From: Sender <fake_sender@example.com>
To: Reader <fake_receiver@example.com>
Subject: Gophercon 2021 is going to be awesome!

Feel free to share my book with others if you're attending.
This recipe can be used to process and parse email information.
`

func main() {
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	printHeaderInfo(m.Header)

	// after printing the header, dump the body to stdout
	if _, err := io.Copy(os.Stdout, m.Body); err != nil {
		log.Fatal(err)
	}
}

/*
$ go run .
To: Reader <fake_receiver@example.com>
From: Sender <fake_sender@example.com>
Subject: Gophercon 2021 is going to be awesome!
Date: 2021-03-18 08:00:00 -0700 -0700
========================================

Feel free to share my book with others if you're attending.
This recipe can be used to process and parse email information.
*/
