package encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// Base64Example demonstrates base64 encoding using the base64 package
func Base64Example() error {
	// base64 is useful for cases where you can't support binary formats
	// it operates on bytes/strings
	var data = "encoding some data!"
	fmt.Println("Source string: ", data)

	// using helper functions and URL encoding
	value := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("With URLEncoding.EncodeToString: ", value)

	// decode the first value
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With URLEncoding.DecodeToString: ", string(decoded))

	return nil
}

// Base64ExampleEncoder shows similar examples with encoders/decoders
func Base64ExampleEncoder() error {
	// using encoder/decoder
	buffer := bytes.Buffer{}

	// encode into the buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)
	var data = "encoding some data!"
	if _, err := encoder.Write([]byte(data)); err != nil {
		return err
	}
	// be sure to close
	if err := encoder.Close(); err != nil {
		return err
	}
	fmt.Println("With NewEncoder and StdEncoding: ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}
	fmt.Println("With NewDecoder and StdEncoding: ", string(results))

	return nil
}
