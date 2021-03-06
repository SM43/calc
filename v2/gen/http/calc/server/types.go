// Code generated by goa v3.2.6, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen github.com/sm43/calc/v2/design

package server

import (
	calc "github.com/sm43/calc/v2/gen/calc"
)

// NewV2AddPayload builds a calc service v2Add endpoint payload.
func NewV2AddPayload(a int, b int) *calc.V2AddPayload {
	v := &calc.V2AddPayload{}
	v.A = a
	v.B = b

	return v
}
