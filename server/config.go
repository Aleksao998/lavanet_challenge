package server

import (
	"net"

	"github.com/hashicorp/go-hclog"
)

// Config is used to parametrize the lavanet_challenge client
type Config struct {
	// NetworkGrpcAddress is network gRPC endpoint
	NetworkGrpcAddress *net.TCPAddr

	// GrpcAddress is gRPC address of lavanet_challenge client
	GrpcAddress *net.TCPAddr

	// logLevel represent a log type [ERROR, INFO, DEBUG]
	LogLevel hclog.Level

	// logFileLocation location of log file
	LogFilePath string
}
