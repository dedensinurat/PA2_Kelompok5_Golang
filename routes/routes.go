package routes

import (
	"API/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up user-related routes
func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
	}
}
