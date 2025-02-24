package controllers

import (
	"API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users)
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	number, err := strconv.ParseUint(id, 10, 32)
	for _, user := range models.Users {
		if user.ID == int(number) {
			c.JSON(http.StatusOK, user)
			return
		} else {
			print(err)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
