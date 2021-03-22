package controllers

// Controller passes state to our handlers
// methods defined in get.go/post.go
type Controller struct {
	storage Storage
}

// New is a Controller 'constructor'
func New(storage Storage) *Controller {
	return &Controller{
		storage: storage,
	}
}

// Payload is our common response
type Payload struct {
	Value string `json:"value"`
}
