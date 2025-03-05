package main

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/pkg/s3storage"
	"SparkGuardBackend/services/orchestrator"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
)

type Server struct {
	orchestrator.UnimplementedOrchestratorServer
}

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Intercepting call: %s", info.FullMethod)

	// Извлекаем метаданные из контекста
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	if len(md["authorization"]) < 1 {
		return nil, errors.New("authorization token is missing")
	}

	// Проверяем токен авторизации
	var runner *db.Runner
	var err error
	if runner, err = db.GetRunnerByToken(md["authorization"][0]); err != nil {
		return nil, errors.New("unauthorized request: invalid token")
	}

	return handler(context.WithValue(ctx, "runner", runner), req)
}

func (s *Server) GetNewTask(ctx context.Context, in *emptypb.Empty) (*orchestrator.GetNewTaskResponse, error) {
	// Пример: Проверка метаданных (авторизация)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		authToken := md["authorization"]
		fmt.Printf("Authorization token: %v\n", authToken)
	}

	// Пример создания задачи
	task := &orchestrator.Task{
		ID:     1,
		WorkID: 101,
		Tag:    "example",
		Status: 0,
	}

	return &orchestrator.GetNewTaskResponse{Task: task}, nil
}

func main() {
	if err := s3storage.Connect(os.Getenv("S3_ENDPOINT"), os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET")); err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":666")
	if err != nil {
		log.Fatalf("Failed to listen on port 666: %v", err)
	}

	// Создаем gRPC-сервер
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
	)

	// Регистрируем наш сервис
	orchestrator.RegisterOrchestratorServer(grpcServer, &Server{})

	log.Println("gRPC server is running on port :666")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
