package client

import (
	"fmt"
	"net/http"
)

// Controller embeds an http.Client and uses it internally
type Controller struct {
	*http.Client
}

// Get with a controller object
func (c *Controller) Get(url string) error {
	resp, err := c.Client.Get(url)
	if err != nil {
		return err
	}
	fmt.Printf("HTTP Status %s: %d\n", url, resp.StatusCode)
	return nil
}
