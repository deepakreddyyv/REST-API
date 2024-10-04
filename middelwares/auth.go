package middelwares

import (
	"net/http"

	"deepak.com/web_rest/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	var token string = ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	userId, err := utils.Verify(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
		return
	}

	ctx.Set("userId", userId)

	ctx.Next()
}
