package adoptions

import (
	"SparkGuardBackend/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupControllers(r *gin.Engine) {
	r.GET("/adoptions/:work_id", GetAdoptionsByWork)
	r.GET("/adoptions/:work_id/related", GetRelatedAdoptions)
}

// GetAdoptionsByWork retrieves all adoptions for a specific work
// @Summary Get adoptions by work
// @Description Get the list of all adoptions related to a specific work
// @Security		ApiKeyAuth
// @Tags adoptions
// @Param work_id path int true "Work ID"
// @Produce json
// @Success 200 {object} []db.Adoption
// @Failure 400
// @Failure 500
// @Router /adoptions/{work_id} [get]
func GetAdoptionsByWork(c *gin.Context) {
	workID, err := strconv.Atoi(c.Param("work_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid work ID"})
		return
	}

	adoptions, err := db.GetAdoptionsByWork(uint(workID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve adoptions"})
		return
	}

	c.JSON(http.StatusOK, adoptions)
}

// GetRelatedAdoptions retrieves all related adoptions (recursive)
// @Summary Get related adoptions
// @Description Get all related adoptions recursively for a given work
// @Security		ApiKeyAuth
// @Tags adoptions
// @Param work_id path int true "Work ID"
// @Produce json
// @Success 200 {object} []db.Adoption
// @Failure 400
// @Failure 500
// @Router /adoptions/{work_id}/related [get]
func GetRelatedAdoptions(c *gin.Context) {
	workID, err := strconv.Atoi(c.Param("work_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid work ID"})
		return
	}

	relatedAdoptions, err := db.GetRelatedAdoptions(uint(workID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve related adoptions"})
		return
	}

	c.JSON(http.StatusOK, relatedAdoptions)
}
