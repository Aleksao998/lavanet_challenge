package server

import (
	"fmt"

	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/Aleksao998/lavanet_challenge/proxy"
	"github.com/hashicorp/go-hclog"
)

// Server is the central manager of the lavanet_challenge
type Server struct {
	logger hclog.Logger

	// config server config
	config *Config

	forwardProxy *proxy.ForwardProxy
}

// NewServer creates a new lavanet_challenge server, using the passed in configuration
func NewServer(config *Config) (*Server, error) {
	// initialize logger
	logger, err := helper.NewLoggerFromConfig(config.LogLevel, config.LogFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not setup new logger instance, %w", err)
	}

	// initialize server
	server := &Server{
		logger: logger,
		config: config,
	}

	// initialize forward proxy
	forwardProxy := proxy.NewForwardProxy(
		server.logger,
		server.config.GrpcAddress,
		server.config.NetworkGrpcAddress,
	)

	// assign forward proxy to the server
	server.forwardProxy = forwardProxy

	// setup and start forward proxy
	if err := forwardProxy.Start(); err != nil {
		return nil, err
	}

	return server, nil
}

// Close closes the lavanet_challenge server
func (s *Server) Close() {
	s.logger.Debug("Closing server")

	s.forwardProxy.Close()
}
