package events

import (
	"SparkGuardBackend/controllers/basic"
	"SparkGuardBackend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Retrieve all events
// @Description Get a list of all events
// @Tags Events
// @Produce json
// @Success 200 {array} db.Event
// @Failure 500 {object} basic.DefaultErrorResponse
// @Router /event/ [get]
func getEvents(c *gin.Context) {
	events, err := db.GetEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Can't get events",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var request CreateEventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err := db.CreateEvent(&request.Event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to create event",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, request.Event)
}

func getEvent(c *gin.Context) {
	var request GetEventRequest
	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid URI parameter",
			Error:   err.Error(),
		})
		return
	}

	event, err := db.GetEvent(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to retrieve event",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

func editEvent(c *gin.Context) {
	var request EditEventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid URI parameter",
			Error:   err.Error(),
		})
		return
	}

	request.Event.ID = request.ID

	err := db.EditEvent(&request.Event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to update event",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, request.Event)
}

func deleteEvent(c *gin.Context) {
	var request DeleteEventRequest
	if err := c.BindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.DefaultErrorResponse{
			Message: "Invalid URI parameter",
			Error:   err.Error(),
		})
		return
	}

	err := db.DeleteEvent(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, basic.DefaultErrorResponse{
			Message: "Failed to delete event",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func SetupControllers(r *gin.Engine) {
	students := r.Group("/event")
	{
		students.GET("/", getEvents)
		students.POST("/", createEvent)
		students.GET("/:id", getEvent)
		students.PATCH("/:id", editEvent)
		students.DELETE("/:id", deleteEvent)
	}
}
