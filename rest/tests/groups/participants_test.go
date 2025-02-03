package groups

import (
	"SparkGuardBackend/controllers"
	"SparkGuardBackend/db"
	"SparkGuardBackend/tests/students"
	"SparkGuardBackend/tests/users"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	db.DeleteAll()

	router := controllers.SetupRouter()

	user, _, err := users.CreateUser(router, "Test1", "test@student.ru", 1)
	if err != nil {
		t.Error("Failed to create student:", err)
		return
	}

	student, err := students.CreateStudent(router, "Test1", "test@student.ru", nil)
	if err != nil {
		t.Error("Failed to create student:", err)
		return
	}

	// Test CreateGroup
	group, err := CreateGroup(router, "Test Group")
	if err != nil {
		t.Error("Failed to create group:", err)
		return
	}

	// Test AddUserToGroup and RemoveUserFromGroup
	err = AddUserToGroup(router, group.ID, user.ID)
	if err != nil {
		t.Error("Failed to add user to group:", err)
	}

	err = RemoveUserFromGroup(router, group.ID, user.ID)
	if err != nil {
		t.Error("Failed to remove user from group:", err)
	}

	// Test AddStudentToGroup and RemoveStudentFromGroup
	err = AddStudentToGroup(router, group.ID, student.ID)
	if err != nil {
		t.Error("Failed to add student to group:", err)
	}

	err = RemoveStudentFromGroup(router, group.ID, student.ID)
	if err != nil {
		t.Error("Failed to remove student from group:", err)
	}
}
