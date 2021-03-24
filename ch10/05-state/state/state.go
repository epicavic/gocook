package state

type op string

const (
	Add      = "add"  // Add values
	Subtract = "sub"  // Subtract values
	Multiply = "mult" // Multiply values
	Divide   = "div"  // Divide values
)

// WorkRequest perform an op on two values
type WorkRequest struct {
	Op op
	V1 int64
	V2 int64
}

// WorkResponse returns the result and any errors
type WorkResponse struct {
	Wr     WorkRequest
	Result int64
}
