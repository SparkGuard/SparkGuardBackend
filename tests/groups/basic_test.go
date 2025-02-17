package groups

import (
	"SparkGuardBackend/cmd/rest/controllers"
	"SparkGuardBackend/db"
	"testing"
)

func TestGroups(t *testing.T) {
	db.DeleteAll()

	router := controllers.SetupRouter()

	// Test CreateGroup
	group, err := CreateGroup(router, "Test Group")
	if err != nil {
		t.Error("Failed to create group:", err)
		return
	}

	// Test GetGroup
	getGroup, err := GetGroup(router, group.ID)

	if err != nil {
		t.Error("Failed to get group:", err)
		return
	}

	if getGroup.Name != group.Name {
		t.Errorf("Expected group name '%s', got '%s'", group.Name, getGroup.Name)
	}

	// Test EditGroup
	group.Name = "Edited Test Group"
	err = EditGroup(router, group)
	if err != nil {
		t.Error("Failed to edit group:", err)
		return
	}

	editedGroup, err := GetGroup(router, group.ID)

	if err != nil {
		t.Error("Failed to get edited group:", err)
		return
	}

	if editedGroup.Name != group.Name {
		t.Errorf("Expected group name '%s', got '%s'", group.Name, editedGroup.Name)
	}
}
