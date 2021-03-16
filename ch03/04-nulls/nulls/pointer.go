package nulls

import (
	"encoding/json"
	"fmt"
)

// PersonPointer is the same as Person, but uses Age pointer to int
// Switching from a value to a pointer is a quick way to express
// null values when marshaling and unmarshaling
type PersonPointer struct {
	Name string `json:"name"`
	Age  *int   `json:"age,omitempty"`
}

// PointerEncoding shows methods for dealing with nil/omitted values
func PointerEncoding() error {
	// note that no age = nil age
	e := PersonPointer{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("jsonBlob: Pointer Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("jsonBlob: Pointer Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("fulljsonBlob: Pointer Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("fulljsonBlob: Pointer Marshal, with age = 0:", string(value))

	return nil
}
