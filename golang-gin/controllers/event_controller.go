package routes

import (
	"golang-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Get all events
// @Description  Get all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Event
// @Router       /events [get]
func Index(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"data":    events,
	})
}

// @Summary      Get an event
// @Description  Get an event
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  models.Event
// @Router       /events/{id} [get]
func Show(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"data":    *event,
	})
}

// @Summary      Create an event
// @Description  Create an event
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event   body      models.Event  true  "Event"
// @Success      200  {object}  models.Event
// @Router       /events [post]
func Store(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	event.ID = len(events) + 1
	event.UserId = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"data":    event,
	})
}
