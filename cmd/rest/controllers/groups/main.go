package groups

import (
	"SparkGuardBackend/cmd/rest/middleware"
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupControllers sets up the routes for the groups controller
func SetupControllers(r *gin.Engine) {
	r.GET("/groups", getGroups)
	r.POST("/groups", middleware.TeacherMiddleware, createGroup)
	r.GET("/groups/:id", getGroup)
	r.PATCH("/groups/:id", middleware.TeacherMiddleware, editGroup)
	r.DELETE("/groups/:id", middleware.TeacherMiddleware, deleteGroup)

	r.POST("/groups/:id/students", middleware.TeacherMiddleware, addStudentToGroup)
	r.DELETE("/groups/:id/students", middleware.TeacherMiddleware, removeStudentFromGroup)

	r.POST("/groups/:id/users", middleware.TeacherMiddleware, addUserToGroup)
	r.DELETE("/groups/:id/users", middleware.TeacherMiddleware, removeUserFromGroup)
}

// getGroups godoc
// @Summary      Get all groups
// @Description  Retrieve a list of all groups
// @Security		ApiKeyAuth
// @Tags         groups
// @Produce      json
// @Success      200  {array}  db.Group
// @Failure      500  {object}  gin.H
// @Router       /groups [get]
func getGroups(c *gin.Context) {
	groups, err := db.GetGroups()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, groups)
}

// createGroup godoc
// @Summary      Create a group
// @Description  Create a new group in the system
// @Security		ApiKeyAuth
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        body  body      CreateGroupRequest  true  "Group data"
// @Success      201  {object}  db.Group
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /groups [post]
func createGroup(c *gin.Context) {
	var request CreateGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var group = db.Group{
		Name: request.Name,
	}

	err := db.CreateGroup(&group)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// getGroup godoc
// @Summary      Get group by ID
// @Description  Retrieve a group's details by its ID
// @Security		ApiKeyAuth
// @Tags         groups
// @Produce      json
// @Param        id   path      int  true  "Group ID"
// @Success      200  {object}  db.Group
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /groups/{id} [get]
func getGroup(c *gin.Context) {
	var request GetGroupRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := db.GetGroup(request.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

// editGroup godoc
// @Summary      Edit group
// @Description  Update details of an existing group
// @Security		ApiKeyAuth
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        id    path      int               true  "Group ID"
// @Param        body  body      EditGroupRequest  true  "Group data"
// @Success      200  {object}  db.Group
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /groups/{id} [patch]
func editGroup(c *gin.Context) {
	var request EditGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.Group.ID = request.ID

	err := db.EditGroup(&request.Group)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request.Group)
}

// deleteGroup godoc
// @Summary      Delete group
// @Description  Delete a group by its ID
// @Security		ApiKeyAuth
// @Tags         groups
// @Produce      json
// @Param        id   path      int  true  "Group ID"
// @Success      200  {object}  gin.H
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /groups/{id} [delete]
func deleteGroup(c *gin.Context) {
	var request DeleteGroupRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DeleteGroup(request.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
