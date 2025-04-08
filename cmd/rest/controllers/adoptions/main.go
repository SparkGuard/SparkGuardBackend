package adoptions

import (
	"SparkGuardBackend/cmd/rest/controllers/basic"
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/pkg/s3storage"
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func SetupControllers(r *gin.Engine) {
	r.GET("/adoptions/:work_id", GetAdoptionsByWork)
	r.GET("/adoptions/:work_id/related", GetRelatedAdoptions)
	r.GET("/adoptions/get/:adoption_id", GetAdoptionSegment)
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

	adoptions, err := db.GetWorkAdoptions(uint(workID))
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

// GetAdoptionSegment retrieves a specific file segment identified by an adoption record
// @Summary Get adoption file segment
// @Description Downloads the work archive, extracts the specific file, and returns the segment identified by the adoption record.
// @Security		ApiKeyAuth
// @Tags adoptions
// @Param adoption_id path int true "Adoption ID"
// @Produce application/octet-stream
// @Success 200 {string} byte "The requested file segment"
// @Failure 400 {object} basic.DefaultErrorResponse "Invalid adoption ID or missing segment info in adoption"
// @Failure 404 {object} basic.DefaultErrorResponse "Adoption, Work, S3 file, or file within archive not found, or segment out of bounds"
// @Failure 500 {object} basic.DefaultErrorResponse "Internal server error (S3 download, zip extraction, file reading)"
// @Router /adoptions/get/{adoption_id} [get]
func GetAdoptionSegment(c *gin.Context) {
	adoptionIDStr := c.Param("adoption_id")
	adoptionID, err := strconv.ParseUint(adoptionIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid adoption ID format"})
		return
	}

	adoption, err := db.GetAdoption(uint(adoptionID))
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, basic.DefaultErrorResponse{Error: fmt.Sprintf("Adoption with ID %d not found", adoptionID)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve adoption record: " + err.Error()})
		}
		return
	}

	if adoption.Path == nil || adoption.PartOffset == nil || adoption.PartSize == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Adoption record is missing required path or segment information"})
		return
	}
	if *adoption.PartSize == 0 {
		c.Data(http.StatusOK, "application/octet-stream", []byte{})
		return
	}

	work, err := db.GetWork(uint(adoption.WorkID))
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Work with ID %d not found", adoption.WorkID)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve work record: " + err.Error()})
		}
		return
	}

	s3Key := fmt.Sprintf("./%d/%d.zip", work.EventID, work.ID)

	zipDataBuffer, err := s3storage.DownloadFileToMemory(s3Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download archive from storage: " + err.Error()})
		return
	}

	zipReader, err := zip.NewReader(bytes.NewReader(zipDataBuffer.Bytes()), int64(zipDataBuffer.Len()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open downloaded archive: " + err.Error()})
		return
	}

	var segmentData []byte
	var fileFound bool

	for _, file := range zipReader.File {
		if file.Name == *adoption.Path {
			fileFound = true
			if file.FileInfo().IsDir() {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Path '%s' points to a directory, not a file", *adoption.Path)})
				return
			}

			rc, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open file '%s' within archive: %s", *adoption.Path, err.Error())})
				return
			}
			defer rc.Close()

			fileSize := int64(file.UncompressedSize64)
			offset := int64(*adoption.PartOffset)
			size := int64(*adoption.PartSize)

			if offset < 0 || size <= 0 || offset >= fileSize {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Invalid segment bounds for file '%s': offset %d, size %d (file size %d)", *adoption.Path, offset, size, fileSize)})
				return
			}

			if offset+size > fileSize {
				size = fileSize - offset
			}

			_, err = rc.(io.Seeker).Seek(offset, io.SeekStart)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to seek in file '%s': %s", *adoption.Path, err.Error())})
				return
			}

			segmentData = make([]byte, size)
			_, err = io.ReadFull(rc, segmentData)
			if err != nil && err != io.ErrUnexpectedEOF {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read segment from file '%s': %s", *adoption.Path, err.Error())})
				return
			}

			break
		}
	}

	if !fileFound {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("File '%s' not found within the archive %s", *adoption.Path, s3Key)})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", segmentData)
}
