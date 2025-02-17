package runner

import "SparkGuardBackend/internal/db"

type CreateRunnerRequest struct {
	Name string `json:"name" binding:"required"`
	Tag  string `json:"tag" binding:"required"`
}

type CreateRunnerResponse struct {
	db.Runner
	Token string `json:"token"`
}

type EditRunnerRequest struct {
	ID uint `uri:"id" binding:"required"`
	db.Runner
}
