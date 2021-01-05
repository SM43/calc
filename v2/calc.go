package v2

import (
	"context"
	"log"

	"github.com/sm43/calc/add"
	calc "github.com/sm43/calc/v2/gen/calc"
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

// V2Add implements v2Add.
func (s *calcsrvc) V2Add(ctx context.Context, p *calc.V2AddPayload) (res int, err error) {
	s.logger.Print("calc.v2Add")

	// Logic would be in a seperate package and reused in multiple versions
	req := add.Request{
		A: p.A,
		B: p.B,
	}
	return req.Add()
}
