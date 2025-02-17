package events

import (
	"SparkGuardBackend/cmd/rest/controllers/events"
	"SparkGuardBackend/db"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"time"
)

func CreateEvent(router *gin.Engine, name, description string, date time.Time, groupID uint) (*db.Event, error) {
	event := events.CreateEventRequest{
		Event: db.Event{
			Name:        name,
			Description: description,
			Date:        date,
			GroupID:     groupID,
		},
	}

	body, _ := json.Marshal(event)

	req, _ := http.NewRequest(http.MethodPost, "/event/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		return nil, fmt.Errorf("expected status code %d, got %d. Body: %s", http.StatusCreated, w.Code, w.Body.String())
	}

	var createdEvent db.Event
	err := json.Unmarshal(w.Body.Bytes(), &createdEvent)

	if err != nil {
		return nil, err
	}

	return &createdEvent, nil
}

func GetEvent(router *gin.Engine, id uint) (*db.Event, error) {

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/event/%d", id), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		if w.Code == http.StatusNotFound {
			return nil, db.ErrNotFound
		}

		return nil, fmt.Errorf("expected status code %d, got %d with body %s", http.StatusOK, w.Code, w.Body.String())
	}

	var event db.Event
	err := json.Unmarshal(w.Body.Bytes(), &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func EditEvent(router *gin.Engine, event *db.Event) error {
	editEventRequest := events.EditEventRequest{
		ID:    event.ID,
		Event: *event,
	}

	body, err := json.Marshal(editEventRequest)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("/event/%d", event.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		return fmt.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	return nil
}

func DeleteEvent(router *gin.Engine, id uint) error {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/event/%d", id), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		return fmt.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	return nil
}
