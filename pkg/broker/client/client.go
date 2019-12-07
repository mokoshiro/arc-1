package client

import (
	"context"

	"github.com/Bo0km4n/arc/internal/dialer"
	"github.com/Bo0km4n/arc/pkg/broker/api/proto"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

type Client interface {
	proto.TrackerClient
	Close() error
}

type client struct {
	proto.TrackerClient
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (Client, error) {
	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	}
	options = append(options, opts...)
	conn, err := dialer.DialContext(ctx, addr, options...)
	if err != nil {
		return nil, err
	}
	return &client{
		TrackerClient: proto.NewTrackerClient(conn),
		conn:          conn,
	}, nil
}

func (c *client) Close() error {
	return c.conn.Close()
}
