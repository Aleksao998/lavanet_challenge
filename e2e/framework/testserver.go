package framework

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/Aleksao998/lavanet_challenge/command"
	"github.com/Aleksao998/lavanet_challenge/command/helper"
	serverCommand "github.com/Aleksao998/lavanet_challenge/command/server"
	"github.com/Aleksao998/lavanet_challenge/server"
	"github.com/hashicorp/go-hclog"
)

const (
	initialPort = 12000
	binaryName  = "../build/lavanet_challenge"
)

const (
	OsmosisMainnetGrpcRaw = "grpc.osmosis.zone:9090"
	localhost             = "localhost"
)

type TestServer struct {
	t *testing.T

	// Config is server config
	Config *server.Config

	// port is reserved port for grpc listener
	Port *ReservedPort

	cmd *exec.Cmd
}

func NewTestServer(t *testing.T) *TestServer {
	t.Helper()

	port := FindAvailablePort(initialPort, initialPort+10000)
	// generate networkGrpcAddress from raw
	network, err := helper.ResolveAddr(
		OsmosisMainnetGrpcRaw,
		command.OsmosisMainnetGrpcEndpoint,
	)

	host := localhost + ":" + port.Port()

	serverGrpc, err := helper.ResolveAddr(
		host,
		command.LocalHostBinding,
	)

	if err != nil {
		t.Fatal(err)
	}

	config := server.Config{
		NetworkGrpcAddress: network,
		GrpcAddress:        serverGrpc,
		LogLevel:           hclog.LevelFromString("INFO"),
		LogFilePath:        "",
	}

	return &TestServer{
		t:      t,
		Config: &config,
		Port:   port,
	}
}

func (t *TestServer) Stop() {
	if t.cmd != nil {
		if err := t.cmd.Process.Kill(); err != nil {
			t.t.Error(err)
		}
	}
}

func (t *TestServer) ReleaseReservedPorts() {
	if err := t.Port.Close(); err != nil {
		t.t.Error(err)
	}

	t.Port = nil
}

func (t *TestServer) Start(ctx context.Context) error {
	serverCmd := serverCommand.GetCommand()
	args := []string{
		serverCmd.Use,
		"--network-grpc-address", t.networkGrpc(),
		"--log-to", "",
		"--log-level", "INFO",
		"--grpc-address", t.localGrpc(),
	}

	t.ReleaseReservedPorts()

	t.cmd = exec.Command(binaryName, args...)

	stdout := io.Writer(os.Stdout)
	t.cmd.Stdout = stdout
	t.cmd.Stderr = stdout

	if err := t.cmd.Start(); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	return nil
}

func (t *TestServer) networkGrpc() string {
	return t.Config.NetworkGrpcAddress.String()
}

func (t *TestServer) localGrpc() string {
	return t.Config.GrpcAddress.String()
}
