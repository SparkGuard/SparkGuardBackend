package groups

import (
	"SparkGuardBackend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupControllers sets up the routes for the groups controller
func SetupControllers(r *gin.Engine) {
	r.GET("/groups", getGroups)
	r.POST("/groups", createGroup)
	r.GET("/groups/:id", getGroup)
	r.PATCH("/groups/:id", editGroup)
	r.DELETE("/groups/:id", deleteGroup)

	r.POST("/groups/:id/students", addStudentToGroup)
	r.DELETE("/groups/:id/students", removeStudentFromGroup)

	r.POST("/groups/:id/users", addUserToGroup)
	r.DELETE("/groups/:id/users", removeUserFromGroup)
}

func getGroups(c *gin.Context) {
	groups, err := db.GetGroups()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, groups)
}

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
