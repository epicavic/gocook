package state

import (
	"context"
)

// Processor routes work to Process
func Processor(ctx context.Context, in chan WorkRequest, out chan WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			out <- Process(wr)
		}
	}
}

// Process switches on operation type, then does work
func Process(wr WorkRequest) WorkResponse {
	resp := WorkResponse{Wr: wr}

	switch wr.Op {
	case Add:
		resp.Result = wr.V1 + wr.V2
	case Subtract:
		resp.Result = wr.V1 - wr.V2
	case Multiply:
		resp.Result = wr.V1 * wr.V2
	case Divide:
		if wr.V2 == 0 {
			break
		}
		resp.Result = wr.V1 / wr.V2
	}
	return resp
}
