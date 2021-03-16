package nulls

import (
	"encoding/json"
	"fmt"
)

// json that has name but not age
const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

func init() {
	fmt.Printf("jsonBlob: raw value - type: %v - %T\n", jsonBlob, jsonBlob)
	fmt.Printf("fulljsonBlob: raw value - type: %v - %T\n", fulljsonBlob, fulljsonBlob)
	fmt.Println()
}

// Person is a basic struct with age and name fields
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

// BaseEncoding shows encoding and decoding with normal types
func BaseEncoding() error {
	// note that no age = 0 age
	e := Person{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("jsonBlob: Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("jsonBlob: Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("fulljsonBlob: Regular Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("fulljsonBlob: Regular Marshal, with age = 0:", string(value))

	return nil
}
