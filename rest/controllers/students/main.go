package students

import (
	"SparkGuardBackend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all students
// @Description Get all students
// @Tags Students
// @Produce json
// @Success 200 {object} []db.Student
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /students/ [get]
func getStudents(c *gin.Context) {
	students, err := db.GetStudents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// @Summary Get student by ID
// @Description Get student by ID
// @Tags Students
// @Produce json
// @Param id path uint true "Student ID"
// @Success 200 {object} GetStudentResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /students/{id} [get]
func getStudent(c *gin.Context) {
	var request GetStudentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := db.GetStudent(request.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetStudentResponse{*student})
}

// @Summary Create student
// @Description Create student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body CreateStudentRequest true "Student"
// @Success 201 {object} GetStudentResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /students/ [post]
func createStudent(c *gin.Context) {
	var err error
	var request CreateStudentRequest

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.CreateStudent(&request.Student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, GetStudentResponse{request.Student})
}

// @Summary Edit student
// @Description Edit student
// @Tags Students
// @Accept json
// @Produce json
// @Param id path uint true "Student ID"
// @Param student body EditStudentRequest true "Student"
// @Success 200 {object} GetStudentResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /students/{id} [patch]
func editStudent(c *gin.Context) {
	var request EditStudentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.Student.ID = request.ID

	err := db.EditStudent(&request.Student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetStudentResponse{request.Student})
}

func SetupControllers(r *gin.Engine) {
	students := r.Group("/students")
	{
		students.GET("/", getStudents)
		students.POST("/", createStudent)
		students.GET("/:id", getStudent)
		students.PATCH("/:id", editStudent)
		// TODO: students.DELETE("/:id", deleteStudent)
	}
}
