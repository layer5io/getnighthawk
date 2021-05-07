package nighthawk

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/layer5io/meshkit/utils"
	nighthawk_client "github.com/layer5io/nighthawk-go/pkg/proto"
)

// Options argument for customizing the client
type Options struct {
	ServerHost string
	ServerPort int32
}

// Client holds the nighthawk client information
type Client struct {
	Handler    nighthawk_client.NighthawkServiceClient
	connection *grpc.ClientConn
}

// New creates a new instance of the nighthawk client connection
func New(opts Options) (*Client, error) {
	if !utils.TcpCheck(&utils.HostPort{
		Address: opts.ServerHost,
		Port:    opts.ServerPort,
	}, nil) {
		return nil, ErrInvalidEndpoint
	}

	var dialOptions []grpc.DialOption
	dialOptions = append(dialOptions, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", opts.ServerHost, opts.ServerPort), dialOptions...)
	if err != nil {
		return nil, ErrGRPCDial(err)
	}

	return &Client{
		Handler:    nighthawk_client.NewNighthawkServiceClient(conn),
		connection: conn,
	}, nil
}

// Close closes the client connection
func (c *Client) Close() error {
	if c.connection != nil {
		return c.connection.Close()
	}
	return nil
}
