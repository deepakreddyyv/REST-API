package routes

import (
	"net/http"
	"strconv"

	"deepak.com/web_rest/models"
	"github.com/gin-gonic/gin"
)

func regesterEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	tokenUserId := ctx.GetInt64("userId")

	uevents, err := models.GetEvents(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "invalid event id"})
		return
	}

	if err = uevents[0].RegisterEvent(tokenUserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error...pls try again.."})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully registered for an event"})
}

func cancleRegistration(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	tokenUserId := ctx.GetInt64("userId")

	var e models.Events

	e.Id = eventId

	rowsEff, err := e.CancleRegistration(tokenUserId)

	if err != nil && rowsEff == -1{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return  
	}else if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return 	
	} 

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully cancled for an event..."})

}
