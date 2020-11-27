package public

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func NewClient(addr, port string) (GrpcClient, error) {
	conn, err := grpc.Dial(net.JoinHostPort(addr, port), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("can't dial GRPC server: %w", err)
	}
	client := NewGrpcClient(conn)
	return client, nil
}
