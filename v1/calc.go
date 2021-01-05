package calcapi

import (
	"context"
	"log"

	"github.com/sm43/calc/add"
	calc "github.com/sm43/calc/v1/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) V1Add(ctx context.Context, p *calc.V1AddPayload) (res int, err error) {
	s.logger.Print("v1.add")

	// Logic would be in a seperate package and reused in multiple versions
	req := add.Request{
		A: p.A,
		B: p.B,
	}
	return req.Add()
}
