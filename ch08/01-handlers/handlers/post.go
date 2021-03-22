package handlers

import (
	"encoding/json"
	"net/http"
)

// GreetingResponse is the JSON Response that
// GreetingHandler returns
type GreetingResponse struct {
	Payload struct {
		Greet string `json:"greet,omitempty"`
		Name  string `json:"name,omitempty"`
		Error string `json:"error,omitempty"`
	} `json:"payload"`
	Successful bool `json:"successful"`
}

// GreetingHandler returns a GreetingResponse which either has errors
// or a useful payload
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	gr := GreetingResponse{}
	if err := r.ParseForm(); err != nil {
		gr.Payload.Error = "bad request"
		gr.Successful = false
		if payload, err := json.Marshal(gr); err == nil {
			w.Write(payload)
		}
	}
	name := r.FormValue("name")
	greet := r.FormValue("greet")

	gr.Successful = true
	gr.Payload.Name = name
	gr.Payload.Greet = greet
	payload, err := json.Marshal(gr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
