package main

import (
	"API/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupUserRoutes(router)

	router.Run(":8000")

}
