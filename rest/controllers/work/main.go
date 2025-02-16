package work

import (
	"SparkGuardBackend/controllers/basic"
	"SparkGuardBackend/db"
	"SparkGuardBackend/s3storage"
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
}

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
		if err == db.ErrNotFound {
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
		if err == db.ErrNotFound {
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
		if err == db.ErrNotFound {
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

func uploadWork(c *gin.Context) {
	var request UploadWorkRequest

	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request URI parameters",
			Error:   err.Error(),
		})
		return
	}

	if err := s3storage.UploadFileSafe(fmt.Sprintf("./%d.zip", request.ID), c.Request.Body); err != nil {
		if err == s3storage.ErrFileExists {
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
