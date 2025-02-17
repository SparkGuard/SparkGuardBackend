package events

import (
	"SparkGuardBackend/cmd/rest/controllers"
	"SparkGuardBackend/db"
	"SparkGuardBackend/tests/groups"
	"testing"
	"time"
)

func TestEvents(t *testing.T) {
	db.DeleteAll()

	router := controllers.SetupRouter()

	// Create a group for the event
	group, err := groups.CreateGroup(router, "Test Group")
	if err != nil {
		t.Error("Failed to create group:", err)
		return
	}

	// Test CreateEvent
	event, err := CreateEvent(router, "Test Event", "Test Description", time.Now(), group.ID)
	if err != nil {
		t.Error("Failed to create event:", err)
		return
	}

	// Test GetEvent
	getEvent, err := GetEvent(router, event.ID)
	if err != nil {
		t.Error("Failed to get event:", err)
		return
	}

	if getEvent.Name != event.Name {
		t.Errorf("Expected event name '%s', got '%s'", event.Name, getEvent.Name)
	}

	// Test EditEvent
	event.Name = "Edited Test Event"
	err = EditEvent(router, event)
	if err != nil {
		t.Error("Failed to edit event:", err)
		return
	}

	editedEvent, err := GetEvent(router, event.ID)
	if err != nil {
		t.Error("Failed to get edited event:", err)
		return
	}

	if editedEvent.Name != event.Name {
		t.Errorf("Expected event name '%s', got '%s'", event.Name, editedEvent.Name)
	}

	// Test DeleteEvent
	err = DeleteEvent(router, event.ID)
	if err != nil {
		t.Error("Failed to delete event:", err)
		return
	}

	_, err = GetEvent(router, event.ID)

	if err == nil {
		t.Error("Event should not exist")
	}
}
