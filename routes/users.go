package routes

import (
	"net/http"

	"deepak.com/web_rest/models"
	"deepak.com/web_rest/utils"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "created user.."})

}


func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
    
	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{"message": err.Error()})
		return 
	}

	err = user.Login()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return 
	}

	jwt, err := utils.GenerateJwtToken(user)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user unauthorised"})
		return 
	}

    ctx.JSON(http.StatusAccepted, gin.H{"message": "valid credentials", "token": jwt})

	
}
