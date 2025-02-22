package students

import (
	"SparkGuardBackend/cmd/rest/controllers"
	"SparkGuardBackend/internal/db"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	db.DeleteAll()

	router := controllers.SetupRouter()

	student, err := CreateStudent(router, "kerblif", "i@kerblif", nil)

	if err != nil {
		t.Error(err)
		return
	}

	student, err = GetStudent(router, student.ID)

	if err != nil {
		t.Error(err)
		return
	}

	if student.Name != "kerblif" {
		t.Errorf("Expected %s, got %s", "kerblif", student.Name)
		return
	}

	if student.Email != "i@kerblif" {
		t.Errorf("Expected %s, got %s", "i@kerblif", student.Email)
		return
	}

	student.Name = "kerblif2"
	student.Email = "anlazarenko@edu.hse.ru"

	if err = EditStudent(router, student); err != nil {
		t.Error(err)
		return
	}

	if student, err = GetStudent(router, student.ID); err != nil {
		t.Error(err)
		return
	}

	if student.Name != "kerblif2" {
		t.Errorf("Expected %s, got %s", "kerblif2", student.Name)
		return
	}

	if student.Email != "anlazarenko@edu.hse.ru" {
		t.Errorf("Expected %s, got %s", "anlazarenko@edu.hse.ru", student.Email)
		return
	}
}
