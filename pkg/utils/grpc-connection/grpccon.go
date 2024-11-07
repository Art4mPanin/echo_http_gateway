package grpc_connection

import (
	"Http-gateway/internal/config"
	"Http-gateway/pkg/singleton"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func ConnectGrpcService(grpcport int) (*grpc.ClientConn, error) {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")
	address := fmt.Sprintf("%s:%d", cfg.Host, grpcport)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	fmt.Printf("Connecting to grpc server %v\n", grpcport)
	return conn, nil
}
