package users

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Produce json
// @Success 200 {object} []db.User
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /users/ [get]
func getUsers(c *gin.Context) {
	users, err := db.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags Users
// @Produce json
// @Param id path uint true "User ID"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	var request GetUserRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.GetUser(request.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetUserResponse{*user})
}

// @Summary Create user
// @Description Create user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /users/ [post]
func createUser(c *gin.Context) {
	var request CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := db.User{
		Name:        request.Username,
		Email:       request.Email,
		AccessLevel: request.AccessLevel,
	}

	err := db.CreateUser(&user, request.Salt, request.Hash)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, GetUserResponse{user})
}

// @Summary Edit user
// @Description Edit user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint true "User ID"
// @Param user body EditUserRequest true "User"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /users/{id} [patch]
func editUser(c *gin.Context) {
	var request EditUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.User.ID = request.ID

	err := db.EditUser(&request.User)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetUserResponse{request.User})
}

func SetupControllers(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", getUsers)
		users.POST("/", createUser)
		users.GET("/:id", getUser)
		users.PATCH("/:id", editUser)
		// TODO: delete handle
	}
}
