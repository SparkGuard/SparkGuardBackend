package middleware

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/internal/auth"
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	var userID uint
	var err error

	if userID, err = auth.ExtractUserFromToken(c); err != nil {
		c.AbortWithStatusJSON(401, basic.DefaultErrorResponse{
			Message: "unauthorized",
			Error:   err.Error(),
		})
		return
	}

	user, err := db.GetUser(userID)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
	}

	c.Set("user", user)

	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	user := c.MustGet("user").(*db.User)

	if user.AccessLevel != "Admin" {
		c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
	}

	c.Next()
}

func TeacherMiddleware(c *gin.Context) {
	user := c.MustGet("user").(*db.User)

	if user.AccessLevel != "Teacher" && user.AccessLevel != "Admin" {
		c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
	}

	c.Next()
}
