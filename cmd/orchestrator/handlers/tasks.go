package handlers

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/services/orchestrator"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (_ *Server) GetNewTask(ctx context.Context, _ *emptypb.Empty) (*orchestrator.GetNewTaskResponse, error) {
	runner, ok := ctx.Value("runner").(db.Runner)

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "runner information is missing from context")
	}

	task, err := db.GetTaskFromQueueForRunner(runner.Tag)

	if err != nil {
		return nil, err
	}

	work, err := db.GetWork(task.WorkID)

	if err != nil {
		return nil, err
	}

	response := &orchestrator.GetNewTaskResponse{
		Task: &orchestrator.Task{
			ID:      uint64(task.ID),
			EventID: uint64(work.EventID),
			WorkID:  uint64(work.ID),
			Tag:     task.Tag,
			Status:  task.Status,
		},
	}

	return response, nil
}

func (_ *Server) CloseTask(_ context.Context, request *orchestrator.CloseTaskRequest) (*emptypb.Empty, error) {
	for _, id := range request.ID {
		if err := db.CloseTask(uint(id)); err != nil {
			fmt.Printf("Failed to close task with ID %d: %v\n", id, err)
		}
	}

	return nil, nil
}
