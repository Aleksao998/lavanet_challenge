package server

import (
	"fmt"
	"os"

	"github.com/Aleksao998/lavanet_challenge/proxy"
	"github.com/hashicorp/go-hclog"
)

// Server is the central manager of the blockchain client
type Server struct {
	logger hclog.Logger

	// config server config
	config *Config

	forwardProxy *proxy.ForwardProxy
}

// NewServer creates a new lavanet_challenge server, using the passed in configuration
func NewServer(config *Config) (*Server, error) {
	// initialize logger
	logger, err := newLoggerFromConfig(config)
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
		logger,
		server.config.GrpcAddress,
		server.config.NetworkGrpcAddress,
	)

	// assign forward proxy to the server
	server.forwardProxy = forwardProxy

	// setup and start forwardProxy
	if err := forwardProxy.Start(); err != nil {
		return nil, err
	}

	return server, nil
}

// newLoggerFromConfig creates a new logger which logs to a specified file or standard output.
func newLoggerFromConfig(config *Config) (hclog.Logger, error) {
	if config.LogFilePath != "" {
		fileLoggerInstance, err := newFileLogger(config)
		if err != nil {
			return nil, err
		}

		return fileLoggerInstance, nil
	}

	return newCLILogger(config), nil
}

// newFileLogger returns logger instance that writes all logs to a specified file.
func newFileLogger(config *Config) (hclog.Logger, error) {
	logFileWriter, err := os.Create(config.LogFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create log file, %w", err)
	}

	return hclog.New(&hclog.LoggerOptions{
		Name:   "lavanet_challenge",
		Level:  config.LogLevel,
		Output: logFileWriter,
	}), nil
}

// newCLILogger returns logger instance that sends all logs to standard output
func newCLILogger(config *Config) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:  "lavanet_challenge",
		Level: config.LogLevel,
	})
}

// Close closes the lavanet_challenge server
func (s *Server) Close() {
	s.logger.Debug("Closing server")
	s.forwardProxy.Close()
}
