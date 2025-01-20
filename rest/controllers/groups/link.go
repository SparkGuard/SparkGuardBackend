package groups

import (
	"SparkGuardBackend/db"
	"github.com/gin-gonic/gin"
)

func addUserToGroup(c *gin.Context) {
	var request AddUserToGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.AddUserToGroup(request.UserID, request.GroupID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func removeUserFromGroup(c *gin.Context) {
	var request RemoveUserFromGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.RemoveUserFromGroup(request.UserID, request.GroupID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func addStudentToGroup(c *gin.Context) {
	var request AddStudentToGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.AddStudentToGroup(request.StudentID, request.GroupID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func removeStudentFromGroup(c *gin.Context) {
	var request RemoveStudentFromGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.RemoveStudentFromGroup(request.StudentID, request.GroupID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
