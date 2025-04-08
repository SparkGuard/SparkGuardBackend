package work

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/cmd/rest/middleware"
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/internal/repacker"
	"SparkGuardBackend/pkg/s3storage"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SetupControllers(r *gin.Engine) {
	r.GET("/works", getWorks)
	r.POST("/works", createWork)
	r.GET("/works/:id", getWork)
	r.PUT("/works/:id", editWork)
	r.DELETE("/works/:id", middleware.TeacherMiddleware, deleteWork)
	r.PUT("/works/:id/upload", uploadWork)
	r.GET("/works/:id/download", downloadWork)
	r.GET("/works/:id/adoptions/download", downloadWorkAdoptionSegments)
}

// @Summary Retrieves all works
// @Description Fetches and returns a list of all works from the database
// @Security		ApiKeyAuth
// @Tags Works
// @Accept json
// @Produce json
// @Success 200 {array} db.Work
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error"
// @Router /works [get]
func getWorks(c *gin.Context) {
	works, err := db.GetWorks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, works)
}

// @Summary Retrieves a specific work by ID
// @Description Fetches a single work item by its ID
// @Security		ApiKeyAuth
// @Tags Works
// @Accept json
// @Produce json
// @Param id path uint true "Work ID"
// @Success 200 {object} db.Work
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid request URI parameters"
// @Failure 404 {object} basic.DefaultErrorResponse "Work not found"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error"
// @Router /works/{id} [get]
func getWork(c *gin.Context) {
	var request GetWorkRequest

	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request URI parameters",
			Error:   err.Error(),
		})
		return
	}

	work, err := db.GetWork(request.ID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{
				Message: "Work not found",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, work)
}

// @Summary Creates a new work
// @Description Adds a new work item to the database
// @Security		ApiKeyAuth
// @Tags Works
// @Accept json
// @Produce json
// @Param work body CreateWorkRequest true "Work data"
// @Success 201 {object} db.Work
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid request body"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error"
// @Router /works [post]
func createWork(c *gin.Context) {
	var request CreateWorkRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err := db.CreateWork(&request.Work)
	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, request.Work)
}

// @Summary Updates a specific work
// @Description Modifies an existing work item in the database
// @Security		ApiKeyAuth
// @Tags Works
// @Accept json
// @Produce json
// @Param id path uint true "Work ID"
// @Param work body EditWorkRequest true "Work data"
// @Success 200 {object} db.Work
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid request body or URI"
// @Failure 404 {object} basic.DefaultErrorResponse "Work not found"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error"
// @Router /works/{id} [put]
func editWork(c *gin.Context) {
	var request EditWorkRequest

	if err := c.ShouldBindJSON(&request.Work); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request URI parameters",
			Error:   err.Error(),
		})
		return
	}

	// Update the work in the database
	err := db.EditWork(&request.Work)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{
				Message: "Work not found",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, request.Work)
}

// @Summary Deletes a specific work
// @Description Removes a work item from the database by its ID
// @Security		ApiKeyAuth
// @Tags Works
// @Accept json
// @Produce json
// @Param id path uint true "Work ID"
// @Success 200 {object} basic.DefaultSuccessResponse
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid request URI parameters"
// @Failure 404 {object} basic.DefaultErrorResponse "Work not found"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error"
// @Router /works/{id} [delete]
func deleteWork(c *gin.Context) {
	var request DeleteWorkRequest

	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request URI parameters",
			Error:   err.Error(),
		})
		return
	}

	err := db.DeleteWork(request.ID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{
				Message: "Work not found",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, basic.DefaultSuccessResponse{Message: "OK"})
}

// @Summary Uploads a file for a specific work
// @Description Uploads a file for a work item
// @Security		ApiKeyAuth
// @Tags Works
// @Accept application/octet-stream
// @Produce json
// @Param id path uint true "Work ID"
// @Param file formData file true "File to upload"
// @Success 200 {object} basic.DefaultSuccessResponse "File uploaded successfully"
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid request URI parameters"
// @Failure 409 {object} basic.DefaultErrorResponse "File already exists"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error"
// @Router /works/{id}/upload [put]
func uploadWork(c *gin.Context) {
	var request UploadWorkRequest

	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request URI parameters",
			Error:   err.Error(),
		})
		return
	}

	repacked_zip, err := repacker.Repack(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to process the file",
			Error:   err.Error(),
		})
		return
	}

	work, err := db.GetWork(request.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{
			Message: "Work not found",
			Error:   err.Error(),
		})
		return
	}

	if err = s3storage.UploadFileSafe(fmt.Sprintf("./%d/%d.zip", work.EventID, request.ID), repacked_zip); err != nil {
		if errors.Is(err, s3storage.ErrFileExists) {
			c.JSON(http.StatusConflict, basic.DefaultErrorResponse{
				Message: "File already exists",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to upload the file",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, basic.DefaultSuccessResponse{Message: "OK"})
}

// @Summary Downloads a specific work file
// @Description Generates a presigned URL to download a work file
// @Security		ApiKeyAuth
// @Tags Works
// @Produce json
// @Param id path uint true "Work ID"
// @Param redirect query bool false "Redirect flag"
// @Success 302 {string} string "Temporary redirected URL for file download"
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid request URI parameters"
// @Failure 404 {object} basic.DefaultErrorResponse "Work doesn't exist"
// @Failure 500 {object} basic.DefaultErrorResponse "Failed to download the file"
// @Router /works/{id}/download [get]
func downloadWork(c *gin.Context) {
	var err error
	var request DownloadWorkRequest

	if err = c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request query parameters",
			Error:   err.Error(),
		})
		return
	}

	if err = c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request URI parameters",
			Error:   err.Error(),
		})
		return
	}

	work, err := db.GetWork(request.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{
			Message: "Work not found",
			Error:   err.Error(),
		})
		return
	}

	var response DownloadWorkResponse
	if response.URL, err = s3storage.ShareFile(fmt.Sprintf("./%d/%d.zip", work.EventID, request.ID)); err != nil {
		if errors.Is(err, s3storage.ErrFileExists) {
			c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{
				Message: "Work doesn't exist",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to download the file",
			Error:   err.Error(),
		})
		return
	}

	if request.RedirectFlag {
		c.Redirect(http.StatusTemporaryRedirect, response.URL)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// downloadWorkAdoptionSegments downloads a zip archive containing all related adoption segments for a work.
// @Summary Download all related adoption segments
// @Description Retrieves all adoptions related to the specified work (including recursive relations) and packages their file segments into a downloadable zip archive. Each file in the archive contains metadata followed by the segment content.
// @Security		ApiKeyAuth
// @Tags Works
// @Param id path uint true "Work ID"
// @Produce application/zip
// @Success 200 {string} byte "A zip archive containing adoption segments"
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid work ID"
// @Failure 404 {object} basic.DefaultErrorResponse "Work or related adoptions not found"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error during processing or S3 interaction"
// @Router /works/{id}/adoptions/download [get]
func downloadWorkAdoptionSegments(c *gin.Context) {
	workIDStr := c.Param("id")
	workID, err := strconv.ParseUint(workIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid work ID format"})
		return
	}

	relatedAdoptions, err := db.GetRelatedAdoptions(uint(workID))
	if err != nil {
		log.Printf("ERROR: Failed to retrieve related adoptions for work %d: %v", workID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve related adoptions"})
		return
	}

	outputFilename := fmt.Sprintf("work_%d_adoptions.zip", workID)

	if len(relatedAdoptions) == 0 {
		log.Printf("INFO: No related adoptions found for work %d. Returning empty archive.", workID)
		c.Header("Content-Type", "application/zip")
		c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, outputFilename))
		emptyZip := createEmptyZip()
		c.Data(http.StatusOK, "application/zip", emptyZip.Bytes())
		return
	}

	processor := NewSegmentProcessor()

	archiveBuffer, err := createAdoptionsArchive(relatedAdoptions, workID, processor)
	if err != nil {
		log.Printf("ERROR: Failed to create adoptions archive for work %d: %v", workID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate the archive"})
		return
	}

	log.Printf("INFO: Successfully generated archive '%s' for work %d with %d segments.", outputFilename, workID, len(relatedAdoptions))
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, outputFilename))
	c.Data(http.StatusOK, "application/zip", archiveBuffer.Bytes())
}
