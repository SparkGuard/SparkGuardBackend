package tasks

import "SparkGuardBackend/internal/db"

type TaskRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type TaskResponse struct {
	Task db.Task `json:"task"`
}

type WorkTasksRequest struct {
	WorkID uint `uri:"id" binding:"required"`
}
