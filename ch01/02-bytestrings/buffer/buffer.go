package buffer

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some tricks for initializing bytes Buffers
// These buffers implement an io.Reader interface
func toBuffer(rawString string) *bytes.Buffer {

	// We'll start with a string encoded into raw bytes
	// rawBytes := []byte(rawString)

	// There are a number of ways to create a buffer from the raw bytes or from the original string
	// var b = new(bytes.Buffer)
	// b.Write(rawBytes)

	// Alternatively
	// b = bytes.NewBuffer(rawBytes)

	// And avoiding the intial byte array altogether
	// b = bytes.NewBufferString(rawString)

	return bytes.NewBufferString(rawString)
}

// ToString is an example of taking an io.Reader and consuming it all, then returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
