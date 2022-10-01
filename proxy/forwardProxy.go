package proxy

import (
	"net"

	"github.com/hashicorp/go-hclog"
)

// ForwardProxy is a service that listens for gRPC calls
// forwards them to the client and returns response
type ForwardProxy struct {
	// serviceServer gGrpc service
	serviceServer serviceServer

	logger hclog.Logger
}

func NewForwardProxy(
	logger hclog.Logger,
	networkGrpcAddress *net.TCPAddr,
	grpcAddress *net.TCPAddr,
) *ForwardProxy {
	return &ForwardProxy{
		serviceServer: NewServiceServer(
			logger,
			grpcAddress,
			networkGrpcAddress,
		),
		logger: logger.Named("forward-proxy"),
	}
}

// Start starts all forwardProxy service
func (s *ForwardProxy) Start() error {
	return s.serviceServer.Start()
}

// Close closes all forwardProxy services
func (s *ForwardProxy) Close() {
	s.logger.Debug("Closing ForwardProxy service")

	s.serviceServer.Close()
}
