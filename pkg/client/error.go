package nighthawk

import (
	"github.com/layer5io/meshkit/errors"
)

const (
	ErrGRPCDialCode        = "1000"
	ErrInvalidEndpointCode = "1001"
	ErrResponseNilCode     = "1002"
)

var (
	ErrInvalidEndpoint = errors.NewDefault(ErrInvalidEndpointCode, "Endpoint is unavailable or endpoint is unreachable")
	ErrResponseNil     = errors.NewDefault(ErrResponseNilCode, "Response is nil from the generator")
)

func ErrGRPCDial(err error) error {
	return errors.NewDefault(ErrGRPCDialCode, "Error creating nighthawk client", err.Error())
}
