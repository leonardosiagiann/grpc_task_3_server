package main

import (
	"log"
	"net"

	depositService "grpc_server/controllers/account"
	proto "grpc_server/proto/account"

	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterDepositServiceServer(s, &depositService.DepositoService{})

	log.Printf("Listening on port %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
