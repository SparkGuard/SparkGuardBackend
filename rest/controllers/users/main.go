package users

import (
	"SparkGuardBackend/db"
	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Description Get all users
// @Produce json
// @Success 200 {object} []db.User
// @Failure 500 {object} map[string]string
// @Router /users/ [get]
func getUsers(c *gin.Context) {
	users, err := db.GetUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// @Summary Get user by ID
// @Description Get user by ID
// @Produce json
// @Param id path uint true "User ID"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	var request GetUserRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := db.GetUser(request.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, GetUserResponse{*user})
}

// @Summary Create user
// @Description Create user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/ [post]
func createUser(c *gin.Context) {
	var request CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := db.User{
		Name:        request.Username,
		Email:       request.Email,
		AccessLevel: request.AccessLevel,
	}

	err := db.CreateUser(&user, request.Salt, request.Hash)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, GetUserResponse{user})
}

// @Summary Edit user
// @Description Edit user
// @Accept json
// @Produce json
// @Param id path uint true "User ID"
// @Param user body EditUserRequest true "User"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [patch]
func editUser(c *gin.Context) {
	var request EditUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	request.User.ID = request.ID

	err := db.EditUser(&request.User)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, GetUserResponse{request.User})
}

func SetupControllers(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", getUsers)
		users.POST("/", createUser)
		users.GET("/:id", getUser)
		users.PATCH("/:id", editUser)
	}
}
