package public

import (
	"net"

	"google.golang.org/grpc"
)

func NewClient(addr, port string) (GrpcClient, error) {
	conn, err := grpc.Dial(net.JoinHostPort(addr, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := NewGrpcClient(conn)
	return client, nil
}
