package students

import (
	"SparkGuardBackend/internal/db"
)

type GetStudentRequest struct {
	ID uint `uri:"id"`
}

type GetStudentResponse struct {
	db.Student
}

type CreateStudentRequest struct {
	db.Student
}

type EditStudentRequest struct {
	ID uint `uri:"id"`
	db.Student
}
