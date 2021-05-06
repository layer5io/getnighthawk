package nighthawk

import (
	"github.com/layer5io/meshkit/errors"
)

const (
	ErrGRPCDialCode        = "1000"
	ErrInvalidEndpointCode = "1001"
)

var (
	ErrInvalidEndpoint = errors.NewDefault(ErrInvalidEndpointCode, "Endpoint not reachable")
)

func ErrGRPCDial(err error) error {
	return errors.NewDefault(ErrGRPCDialCode, "Error creating nighthawk client", err.Error())
}
