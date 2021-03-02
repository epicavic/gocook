package interfaces

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Copy copies the data from 'in' to 'out' using standard copy and using a buffer
func Copy(in io.ReadSeeker, out io.Writer) error {
	// works like tee. write to out and stdout
	w := io.MultiWriter(out, os.Stdout)
	var err error

	// copyBuffer with buffer size 32KB (default)
	_, err = io.Copy(w, in)
	if err != nil {
		log.Printf("Copy: standard copy failed: %s\n", err)
		return err
	}

	// io.Copy would move the read position to EOF. so we have to rewind to the start
	in.Seek(0, 0)

	// copyBuffer with buffer size 64 byte
	buf := make([]byte, 64)
	_, err = io.CopyBuffer(w, in, buf)
	if err != nil {
		log.Printf("Copy: standard copy failed: %s\n", err)
		return err
	}

	fmt.Println() // add newline
	return nil
}
