package confformat

import "fmt"

// MarshalAll takes some data stored in structs
// and converts them to the various data formats
func MarshalAll() error {
	t := TOMLData{
		Name: "Name1",
		Age:  20,
	}

	j := JSONData{
		Name: "Name2",
		Age:  30,
	}

	y := YAMLData{
		Name: "Name3",
		Age:  40,
	}

	tomlRes, err := t.Encode()
	if err != nil {
		return err
	}
	fmt.Printf("TOML Marshal:\n%v\n", tomlRes.String())

	yamlRes, err := y.Encode()
	if err != nil {
		return err
	}
	fmt.Printf("YAML Marshal:\n%v\n", yamlRes.String())

	jsonRes, err := j.Encode()
	if err != nil {
		return err
	}
	fmt.Printf("JSON Marshal:\n%v\n", jsonRes.String())

	fmt.Println()
	return nil
}
