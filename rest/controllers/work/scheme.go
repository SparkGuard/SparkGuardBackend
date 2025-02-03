package work

import "time"

type CreateWorkRequest struct {
	Time      time.Time `json:"time"`
	StudentID uint      `json:"student_id"`
	EventID   uint      `json:"event_id"`
}

type GetWorkRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type EditWorkRequest struct {
	ID        uint      `uri:"id" binding:"required"`
	Time      time.Time `json:"time"`
	StudentID uint      `json:"student_id"`
	EventID   uint      `json:"event_id"`
}

type DeleteWorkRequest struct {
	ID uint `uri:"id" binding:"required"`
}
