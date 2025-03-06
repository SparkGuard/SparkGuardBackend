package middleware

import (
	"SparkGuardBackend/internal/db"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func AuthInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
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

	return handler(context.WithValue(ctx, "runner", *runner), req)
}
