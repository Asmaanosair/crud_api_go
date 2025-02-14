package controllers

import (
	"goTask/config"
	"goTask/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Succssfully"})
}
