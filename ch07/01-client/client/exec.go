package client

import (
	"fmt"
	"net/http"
)

// GetStatusCode returns request status
func GetStatusCode(c *http.Client, url string) error {
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	fmt.Printf("HTTP Status %s: %d\n", url, resp.StatusCode)

	return nil
}
