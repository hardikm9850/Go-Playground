package operation

import "fmt"

// OperationRequest represents a request for an operation on operands.
type OperationRequest struct {
	Operation string `json:"operation" validate:"required"`
	Operands  []int  `json:"operands" validate:"required,min=2"`
}

func (r *OperationRequest) String() string {
	return fmt.Sprintf("Operation: %s, Operands: %v", r.Operation, r.Operands)
}
