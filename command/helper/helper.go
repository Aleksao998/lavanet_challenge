package helper

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aleksao998/lavanet_challenge/command"
	"github.com/hashicorp/go-hclog"
)

type ClientCloseResult struct {
	Message string `json:"message"`
}

func (r *ClientCloseResult) GetOutput() string {
	return r.Message
}

// HandleSignals is a helper method for handling signals sent to the console
// Like stop, error, etc.
func HandleSignals(
	closeFn func(),
	outputter command.OutputFormatter,
) error {
	signalCh := getTerminationSignalCh()
	sig := <-signalCh

	closeMessage := fmt.Sprintf("\n[SIGNAL] Caught signal: %v\n", sig)
	closeMessage += "Gracefully shutting down client...\n"

	outputter.SetCommandResult(
		&ClientCloseResult{
			Message: closeMessage,
		},
	)
	outputter.WriteOutput()

	// Call the server close callback
	gracefulCh := make(chan struct{})

	go func() {
		if closeFn != nil {
			closeFn()
		}

		close(gracefulCh)
	}()

	select {
	case <-signalCh:
		return errors.New("shutdown by signal channel")
	case <-time.After(5 * time.Second):
		return errors.New("shutdown by timeout")
	case <-gracefulCh:
		return nil
	}
}

// GetTerminationSignalCh returns a channel to emit signals by ctrl + c
func getTerminationSignalCh() <-chan os.Signal {
	// wait for the user to quit with ctrl-c
	signalCh := make(chan os.Signal, 1)
	signal.Notify(
		signalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	return signalCh
}

// ResolveAddr resolves the passed in TCP address
// The second param is the default ip to bind to, if no ip address is specified
func ResolveAddr(address string, defaultIP command.IPBinding) (*net.TCPAddr, error) {
	addr, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		return nil, fmt.Errorf("failed to parse addr '%s': %w", address, err)
	}

	if addr.IP == nil {
		addr.IP = net.ParseIP(string(defaultIP))
	}

	return addr, nil
}

// newFileLogger returns logger instance that writes all logs to a specified file.
func newFileLogger(logLevel hclog.Level, logFilePath string) (hclog.Logger, error) {
	logFileWriter, err := os.Create(logFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create log file, %w", err)
	}

	return hclog.New(&hclog.LoggerOptions{
		Name:   "lavanet_challenge",
		Level:  logLevel,
		Output: logFileWriter,
	}), nil
}

// newCLILogger returns logger instance that sends all logs to standard output
func newCLILogger(logLevel hclog.Level) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:  "lavanet_challenge",
		Level: logLevel,
	})
}

// newLoggerFromConfig creates a new logger which logs to a specified file or standard output.
func NewLoggerFromConfig(logLevel hclog.Level, logFilePath string) (hclog.Logger, error) {
	if logFilePath != "" {
		fileLoggerInstance, err := newFileLogger(logLevel, logFilePath)
		if err != nil {
			return nil, err
		}

		return fileLoggerInstance, nil
	}

	return newCLILogger(logLevel), nil
}
