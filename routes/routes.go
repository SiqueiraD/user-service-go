package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/meajudaaqui/user-service/controllers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("user-api")
	{
		//ROTAS RELACIONADAS A LISTAGEM DE USUÁRIOS
		users := main.Group("usuarios")
		{
			users.GET("/", controllers.MostrarUsers)
		}

		//ROTAS RELACIONADAS CADASTRO DE USUÁRIOS
		signup := main.Group("novo")
		{
			signup.POST("/", controllers.CriarUser)
		}

		//ROTAS RELACIONADAS A AUTENTICAÇÃO E AUTORIZAÇÃO
		auth := main.Group("login")
		{
			auth.POST("/", controllers.Autenticar)
		}
	}
	router.Use(CORSMiddleware())
	return router
}
