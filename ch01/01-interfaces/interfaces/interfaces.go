package interfaces

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
The Go language provides a number of I/O interfaces
It is best practice to make use of these interfaces wherever possible
rather than passing structures or other types directly

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
type ReadSeeker interface {
	Reader
	Seeker
}
*/

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
