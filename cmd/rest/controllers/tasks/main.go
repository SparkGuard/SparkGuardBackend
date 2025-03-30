package tasks

import (
	"SparkGuardBackend/cmd/rest/middleware"
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupControllers(r *gin.Engine) {
	r.GET("/tasks", middleware.AdminMiddleware, GetTasks)
	r.GET("/tasks/:id", middleware.AdminMiddleware, GetTask)
	r.PUT("/tasks/:id/reset", middleware.AdminMiddleware, ResetTask)
}

// GetTasks retrieves all tasks
// @Summary Get tasks
// @Description Get the list of all tasks
// @Security		ApiKeyAuth
// @Tags tasks
// @Produce json
// @Success 200 {object} []db.Task
// @Failure 500
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	tasks, err := db.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a specific task
// @Summary Get a task
// @Description Get a task by its ID
// @Security		ApiKeyAuth
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} db.Task
// @Failure 400
// @Failure 404
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	var request TaskRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := db.GetTask(request.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// ResetTask resets a task's status to "In queue"
// @Summary Reset a task
// @Description Reset a task's status to "In queue"
// @Security		ApiKeyAuth
// @Tags tasks
// @Accept json
// @Param id path int true "Task ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /tasks/{id}/reset [put]
func ResetTask(c *gin.Context) {
	var request TaskRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := db.ResetTask(request.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset task"})
		return
	}

	c.Status(http.StatusNoContent)
}
