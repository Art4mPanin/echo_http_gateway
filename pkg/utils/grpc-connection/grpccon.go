package grpc_connection

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func ConnectGrpcService(grpcport int) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("localhost:%d", grpcport)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	fmt.Printf("Connecting to grpc server %v\n", grpcport)
	return conn, nil
}
