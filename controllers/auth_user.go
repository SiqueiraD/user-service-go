package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/meajudaaqui/user-service/database"
	"github.com/meajudaaqui/user-service/models"
	"github.com/meajudaaqui/user-service/services"
)

func Autenticar(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User

	dbError := db.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "usuario inexistente",
		})
		return
	}

	if user.Password != services.SHA256_encoder(p.Password) {
		c.JSON(401, gin.H{
			"error": "credenciais invalidas",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.Name, user.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.Admin {
		c.JSON(200, gin.H{
			"token": token,
			"admin": true,
		})
	} else {
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}
