package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/meajudaaqui/user-service/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

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

	return router
}
