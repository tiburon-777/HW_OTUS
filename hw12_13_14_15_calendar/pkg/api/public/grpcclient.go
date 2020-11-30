package public

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func NewClient(ctx context.Context, addr, port string) (GrpcClient, error) {
	conn, err := grpc.DialContext(ctx, net.JoinHostPort(addr, port), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("can't dial GRPC server: %w", err)
	}
	client := NewGrpcClient(conn)
	return client, nil
}
