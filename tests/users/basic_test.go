package users

import (
	"SparkGuardBackend/cmd/rest/controllers"
	"SparkGuardBackend/db"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db.DeleteAll()

	router := controllers.SetupRouter()

	user, _, err := CreateUser(router, "kerblif", "anlazarenko@edu.hse.ru", 1)

	if err != nil {
		t.Error(err)
		return
	}

	user, err = GetUser(router, user.ID)

	if err != nil {
		t.Error(err)
		return
	}

	if user.Name != "kerblif" {
		t.Errorf("Expected %s, got %s", "kerblif", user.Name)
		return
	}

	if user.Email != "anlazarenko@edu.hse.ru" {
		t.Errorf("Expected %s, got %s", "anlazarenko@edu.hse.ru", user.Email)
	}

	if user.AccessLevel != 1 {
		t.Errorf("Expected %d, got %d", 1, user.AccessLevel)
	}

	user.Name = "kerblif2"
	user.Email = "i@kerblif.ru"
	user.AccessLevel = 2

	if err = EditUser(router, user); err != nil {
		t.Error(err)
	}

	if user, err = GetUser(router, user.ID); err != nil {
		t.Error(err)
		return
	}

	if user.Name != "kerblif2" {
		t.Errorf("Expected %s, got %s", "kerblif2", user.Name)
		return
	}

	if user.Email != "i@kerblif.ru" {
		t.Errorf("Expected %s, got %s", "i@kerblif.ru", user.Email)
		return
	}

	if user.AccessLevel != 2 {
		t.Errorf("Expected %d, got %d", 2, user.AccessLevel)
		return
	}
}
