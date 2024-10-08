package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"deepak.com/web_rest/models"
	"github.com/gin-gonic/gin"
)

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Request"})
		return
	}

	events, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, events)

}

func getEvents(ctx *gin.Context) { //GET, POST, PUT, PATCH, DELETE
	events, err := models.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvents(ctx *gin.Context) {
	var events models.Events
	err := ctx.ShouldBindJSON(&events) //binds the request object data with events

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Unable to parse the request"})
		return
	}

	//events.Id = 1
	userId := ctx.GetInt64("userId")
	events.UserId = userId

	err = events.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Unable to parse the request"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created the event successfully"})
}

func updateEvents(ctx *gin.Context) {
	var events models.Events
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	tokenUserId := ctx.GetInt64("userId")

	fmt.Println(tokenUserId)

	/* replace it with getEventsById */
	uevent, err := models.GetEventById(id)

	if err != nil || tokenUserId != uevent.UserId { //if user is unauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user - unauthorized"})
		return
	}

	err = ctx.ShouldBindJSON(&events) //binds the request object data with events

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	events.Id = id

	err = events.UpdateEvents(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "please try again"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "updated the event.."})

}

func deleteEvents(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	tokenUserId := ctx.GetInt64("userId")
	/* replace it with getEventsById */
	uevent, err := models.GetEventById(id)

	if err != nil || tokenUserId != uevent.UserId { //if user is unauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user - unauthorized"})
		return
	}

	err = models.DeleteEvents(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "deleted the event.."})
}
