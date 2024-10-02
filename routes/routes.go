package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvents)
	server.DELETE("/events/:id", deleteEvents)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
