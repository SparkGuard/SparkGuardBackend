package users

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/cmd/rest/middleware"
	"SparkGuardBackend/internal/auth"
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all users
// @Description Get all users
// @Security		ApiKeyAuth
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
// @Security		ApiKeyAuth
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
// @Security		ApiKeyAuth
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

	err := db.CreateUser(&request.User, request.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, GetUserResponse{request.User})
}

// @Summary Login user
// @Description Login user
// @Tags Users
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} basic.DefaultErrorResponse
// @Failure 401 {object} basic.DefaultErrorResponse
// @Router /users/login [post]
func loginUser(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.VerifyLogin(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, basic.DefaultErrorResponse{
			Error:   err.Error(),
			Message: "invalid login or password",
		})
		return
	}

	// Генерация JWT токена
	token, err := auth.GenerateJWT(user.ID, user.Email, user.AccessLevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// Возврат информации о пользователе вместе с токеном
	c.JSON(http.StatusOK, LoginResponse{
		User:  user,
		Token: token,
	})
}

func SetupControllers(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", middleware.AuthMiddleware, getUsers)
		users.POST("/", middleware.AuthMiddleware, middleware.AdminMiddleware, createUser)
		users.GET("/:id", middleware.AuthMiddleware, getUser)
		users.POST("/login", loginUser)
	}
}
