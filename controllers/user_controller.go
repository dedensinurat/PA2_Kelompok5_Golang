package controllers

import (
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users)
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range models.Users {
		if id == string(rune(user.ID)) {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
