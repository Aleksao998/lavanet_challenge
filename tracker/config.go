package tracker

import (
	"net"

	"github.com/hashicorp/go-hclog"
)

// Config is used to parametrize the tracker client
type Config struct {
	// ClientGrpcAddress is network gRPC endpoint
	ClientGrpcAddress *net.TCPAddr

	// LogLevel represent a log type [ERROR, INFO, DEBUG]
	LogLevel hclog.Level

	// LogFilePath location of log file
	LogFilePath string

	// PollingTime is a polling time in seconds
	PollingTime uint64

	// OutputFilePath location of output file
	OutputFilePath string
}
