package routes

import (
	"deepak.com/web_rest/middelwares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	/*
    Add a server group and register a middleware(e.g authenticating..) for the protected routes

	here Group("/") says all routes under this group will have a common prefix path 
	for e.g., Group("/api/v1"), all routes now have a common prefix (/api/v1/events, /api/v1/events/:id).
	*/
	routerGroup := server.Group("/")  
	routerGroup.Use(middelwares.Authenticate)
	routerGroup.POST("/events", createEvents)
	routerGroup.PUT("/events/:id", updateEvents)
	routerGroup.DELETE("/events/:id", deleteEvents)
	routerGroup.POST("/events/:id/register", regesterEvent)
	routerGroup.DELETE("/events/:id/register", cancleRegistration)


	// server.POST("/events", middelwares.Authenticate, createEvents)
	// server.POST("/events", createEvents)
	// server.GET("/events/:id", getEvent)
	// server.PUT("/events/:id", updateEvents)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
