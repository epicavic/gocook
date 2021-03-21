package client

import (
	"crypto/tls"
	"net/http"
)

// Setup configures our client
func Setup(isSecure bool) *http.Client {
	c := new(http.Client)

	// skip TLS verification
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return c
}
