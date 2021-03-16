package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type nullInt64 sql.NullInt64

// PersonNullInt is the same, but uses a sql.NullInt64
type PersonNullInt struct {
	Name string     `json:"name"`
	Age  *nullInt64 `json:"age,omitempty"`
}

// MarshalJSON implements to json encoding for custom type nullInt64
func (v *nullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements from json decoding for custom type nullInt64
func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

// NullEncoding shows an alternative method for dealing with nil/omitted values
func NullEncoding() error {
	e := PersonNullInt{}

	// note that no means an invalid value
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("jsonBlob: nullInt64 Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("jsonBlob: nullInt64 Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("fulljsonBlob: nullInt64 Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("fulljsonBlob: nullInt64 Marshal, with age = 0:", string(value))

	return nil
}
