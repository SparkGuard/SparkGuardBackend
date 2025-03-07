package work

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/pkg/s3storage"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupControllers(r *gin.Engine) {
	r.GET("/works", getWorks)
	r.POST("/works", createWork)
	r.GET("/works/:id", getWork)
	r.PUT("/works/:id", editWork)
	r.DELETE("/works/:id", deleteWork)
	r.PUT("/works/:id/upload", uploadWork)
	r.GET("/works/:id/download", downloadWork)
}

// @Summary Retrieves all works
// @Description Fetches and returns a list of all works from the database
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

	repacked_zip, err := repack(c.Request.Body)

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

	var response DownloadWorkResponse

	if response.URL, err = s3storage.ShareFile(fmt.Sprintf("./%d.zip", request.ID)); err != nil {
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
