package runner

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupControllers(r *gin.Engine) {
	r.GET("/runners", GetRunners)
	r.GET("/runners/:id", GetRunner)
	r.POST("/runners", CreateRunner)
	r.PUT("/runners/:id", EditRunner)
	r.DELETE("/runners/:id", DeleteRunner)
}

// GetRunners retrieves all runners
// @Summary Get runners
// @Description Get the list of all runners
// @Tags runners
// @Produce json
// @Success 200 {object} []db.Runner
// @Failure 500
// @Router /runners [get]
func GetRunners(c *gin.Context) {
	runners, err := db.GetRunners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve runners"})
		return
	}
	c.JSON(http.StatusOK, runners)
}

// GetRunner retrieves a specific runner
// @Summary Get a runner
// @Description Get a runner by its ID
// @Tags runners
// @Accept json
// @Produce json
// @Param id path int true "Runner ID"
// @Success 200 {object} db.Runner
// @Failure 400
// @Failure 404
// @Router /runners/{id} [get]
func GetRunner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid runner ID"})
		return
	}

	runner, err := db.GetRunner(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Runner not found"})
		return
	}

	c.JSON(http.StatusOK, runner)
}

// CreateRunner creates a new runner
// @Summary Create a runner
// @Description Create a new runner
// @Tags runners
// @Accept json
// @Produce json
// @Param runner body CreateRunnerRequest true "Create Runner Request"
// @Success 201 {object} db.Runner
// @Failure 400
// @Failure 500
// @Router /runners [post]
func CreateRunner(c *gin.Context) {
	var req CreateRunnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	runner := &db.Runner{
		Name: req.Name,
		Tag:  req.Tag,
	}

	var response CreateRunnerResponse

	if err := db.CreateRunner(runner); err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Error:   err.Error(),
			Message: "Failed to create runner",
		})
		return
	}

	response.Runner = *runner
	response.Token = runner.Token

	c.JSON(http.StatusCreated, response)
}

// EditRunner updates a runner's information
// @Summary Edit a runner
// @Description Edit an existing runner by ID
// @Tags runners
// @Accept json
// @Param id path int true "Runner ID"
// @Param runner body EditRunnerRequest true "Edit Runner Request"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /runners/{id} [put]
func EditRunner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid runner ID"})
		return
	}

	var req EditRunnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	runner := &db.Runner{
		ID:   uint(id),
		Name: req.Name,
		Tag:  req.Tag,
	}

	if err := db.EditRunner(runner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update runner"})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteRunner deletes a runner by ID
// @Summary Delete a runner
// @Description Delete a runner by its ID
// @Tags runners
// @Param id path int true "Runner ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /runners/{id} [delete]
func DeleteRunner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid runner ID"})
		return
	}

	if err = db.DeleteRunner(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Error:   err.Error(),
			Message: "Failed to delete runner",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
