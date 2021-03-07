package flags

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays type is a slice of int that we'll read a flag into
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

// Set method will be used by the flag package
func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}

	return nil
}
