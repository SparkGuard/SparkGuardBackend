package events

import (
	"SparkGuardBackend/internal/db"
)

type CreateEventRequest struct {
	db.Event
}

type GetEventRequest struct {
	ID uint `uri:"id"`
}

type EditEventRequest struct {
	ID uint `uri:"id" binding:"required"`
	db.Event
}

type DeleteEventRequest struct {
	ID uint `uri:"id" binding:"required"`
}
