package students

import (
	"SparkGuardBackend/cmd/rest/controllers/students"
	"SparkGuardBackend/db"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

func CreateStudent(router *gin.Engine, name string, email string, userID *uint) (res *db.Student, err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/students/", nil)
	req.Header.Set("Content-Type", "application/json")

	student := students.CreateStudentRequest{
		db.Student{
			Name:   name,
			Email:  email,
			UserID: userID,
		},
	}
	studentJson, err := json.Marshal(student)

	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(studentJson))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		err = errors.New(w.Body.String())
		return
	}

	res = new(db.Student)
	err = json.Unmarshal(w.Body.Bytes(), res)

	return
}

func GetStudent(router *gin.Engine, id uint) (student *db.Student, err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", fmt.Sprintf("/students/%d", id), nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	student = new(db.Student)
	err = json.Unmarshal(w.Body.Bytes(), &student)

	return
}

func EditStudent(router *gin.Engine, student *db.Student) (err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("PATCH", fmt.Sprintf("/students/%d", student.ID), nil)
	req.Header.Set("Content-Type", "application/json")

	editStudent := students.EditStudentRequest{
		Student: *student,
	}
	editStudentJson, err := json.Marshal(editStudent)

	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(editStudentJson))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	return
}
