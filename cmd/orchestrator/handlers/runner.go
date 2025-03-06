package handlers

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/services/orchestrator"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (_ *Server) GetRunnerInfo(ctx context.Context, _ *emptypb.Empty) (*orchestrator.GetRunnerInfoResponse, error) {
	runner, ok := ctx.Value("runner").(db.Runner)

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "runner information is missing from context")
	}

	response := &orchestrator.GetRunnerInfoResponse{
		Runner: &orchestrator.Runner{
			ID:   uint64(runner.ID),
			Name: runner.Name,
			Tag:  runner.Tag,
		},
	}

	return response, nil
}
