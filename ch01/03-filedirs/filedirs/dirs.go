package filedirs

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Operate manipulates files and directories
func Operate() error {
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}

	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	value := []byte("hello\n")
	_, err = f.Write(value)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Chdir(".."); err != nil {
		return err
	}

	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	// cleanup, os.RemoveAll can be dangerous if you point at the wrong directory
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}
