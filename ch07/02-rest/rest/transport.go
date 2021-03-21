package rest

import "net/http"

// APITransport does a SetBasicAuth for every request
type APITransport struct {
	*http.Transport    // http.RoundTripper interface
	username, password string
}

// RoundTrip does the basic auth before deferring to the default transport
// https://lanre.wtf/blog/2017/07/24/roundtripper-go/
func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.Transport.RoundTrip(req)
}
