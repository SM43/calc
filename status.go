package calc

import (
	"context"
	"log"

	status "github.com/sm43/calc/gen/status"
)

// status service example implementation.
// The example methods log the requests and return zero values.
type statussrvc struct {
	logger *log.Logger
}

// NewStatus returns the status service implementation.
func NewStatus(logger *log.Logger) status.Service {
	return &statussrvc{logger}
}

// Status implements status.
func (s *statussrvc) Status(ctx context.Context) (res string, err error) {
	s.logger.Print("status.status")
	return "OK", nil
}
