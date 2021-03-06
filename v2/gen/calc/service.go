// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen github.com/sm43/calc/v2/design

package calc

import (
	"context"
)

// The calc service performs operations on numbers.
type Service interface {
	// V2Add implements v2Add.
	V2Add(context.Context, *V2AddPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"v2Add"}

// V2AddPayload is the payload type of the calc service v2Add method.
type V2AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}
