package main

import (
	"SparkGuardBackend/cmd/orchestrator/handlers"
	"SparkGuardBackend/cmd/orchestrator/middleware"
	"SparkGuardBackend/services/orchestrator"
	"google.golang.org/grpc"
	"log"
	"net"
)

func serve(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)

	orchestrator.RegisterOrchestratorServer(grpcServer, &handlers.Server{})

	log.Printf("gRPC server is running on %s\n", addr)
	return grpcServer.Serve(listener)
}
