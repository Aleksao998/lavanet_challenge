package server

import (
	"github.com/Aleksao998/lavanet_challenge/server"
	"github.com/hashicorp/go-hclog"
)

const (
	network             = "network"
	port                = "port"
	logFileLocationFlag = "log-to"
	logLevelFlag        = "log-level"
)

var (
	params = &serverParams{}
)

type serverParams struct {
	// network gRPC endpoint
	network string

	// gRPC server port
	port uint64

	// logLevel represent a log type [ERROR, INFO, DEBUG]
	logLevel string

	// logFileLocation location of log file
	logFileLocation string
}

func (p *serverParams) generateConfig() *server.Config {
	return &server.Config{
		Network:     p.network,
		Port:        p.port,
		LogLevel:    hclog.LevelFromString(p.logLevel),
		LogFilePath: p.logFileLocation,
	}
}
