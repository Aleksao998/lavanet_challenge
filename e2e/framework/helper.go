package framework

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	DefaultTimeout = time.Minute
)

// ReservedPort keeps available port until use
type ReservedPort struct {
	port     int
	listener net.Listener
	isClosed bool
}

func (p *ReservedPort) Port() string {
	return strconv.Itoa(p.port)
}

func FindAvailablePort(from, to int) *ReservedPort {
	for port := from; port < to; port++ {
		addr := fmt.Sprintf("localhost:%d", port)
		if l, err := net.Listen("tcp", addr); err == nil {
			return &ReservedPort{port: port, listener: l}
		}
	}

	return nil
}

func (p *ReservedPort) Close() error {
	if p.isClosed {
		return nil
	}

	err := p.listener.Close()
	p.isClosed = true

	return err
}

func NewTestServers(t *testing.T, num int) []*TestServer {
	t.Helper()

	srvs := make([]*TestServer, 0, num)

	t.Cleanup(func() {
		for _, srv := range srvs {
			srv.Stop()
		}
	})

	for i := 0; i < num; i++ {
		srv := NewTestServer(t)
		srvs = append(srvs, srv)
	}

	var wg sync.WaitGroup

	for _, srv := range srvs {
		wg.Add(1)

		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
			defer cancel()

			if err := srv.Start(ctx); err != nil {
				t.Fatal("server failed to start", err)
			}
		}()
	}

	wg.Wait()

	return srvs
}

type v1Beta1Client struct {
	// Client represents network client
	Client tendermintv1beta1.ServiceClient

	// connection is grpc client connection
	connection *grpc.ClientConn
}

func NewV1Beta1Client(grpcAddress string) (*v1Beta1Client, error) {
	conn, err := grpc.Dial(
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &v1Beta1Client{
		Client:     tendermintv1beta1.NewServiceClient(conn),
		connection: conn,
	}, nil
}

// Close closes network client
func (s *v1Beta1Client) Close() {
	s.connection.Close()
}
