package work

import (
	"SparkGuardBackend/internal/db"
)

type CreateWorkRequest struct {
	db.Work
}

type GetWorkRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type EditWorkRequest struct {
	ID uint `json:"-" uri:"id"`
	db.Work
}

type DeleteWorkRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type UploadWorkRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type DownloadWorkRequest struct {
	ID           uint `uri:"id"`
	RedirectFlag bool `form:"redirect"`
}

type DownloadWorkResponse struct {
	URL string `json:"url"`
}
