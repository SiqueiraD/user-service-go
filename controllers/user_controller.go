package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/meajudaaqui/user-service/database"
	"github.com/meajudaaqui/user-service/models"
	"github.com/meajudaaqui/user-service/services"
)

func CriarUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256_encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "nao foi possivel criar o usuario: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func MostrarUsers(c *gin.Context) {
	db := database.GetDatabase()

	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "nao foi possivel listar os usuarios: " + err.Error(),
		})
		return
	}

	c.JSON(200, users)
}
