package pool

import "errors"

type op string

const (
	Hash    = "encrypt" // Hash is the bcrypt enctypt
	Compare = "decrypt" // Compare is bcrypt compare
)

// WorkRequest is a worker req
type WorkRequest struct {
	Op      op     // work type
	Text    []byte // work data
	Compare []byte // work data compare
}

// WorkResponse is a worker resp
type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

// Process dispatches work to the worker pool channel
func Process(wr WorkRequest) WorkResponse {
	switch wr.Op {
	case Hash:
		return hashWork(wr)
	case Compare:
		return compareWork(wr)
	default:
		return WorkResponse{Err: errors.New("unsupported operation")}
	}
}
