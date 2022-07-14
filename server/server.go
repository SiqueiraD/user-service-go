package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/meajudaaqui/user-service/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port if not specified
	}
	return Server{
		port:   port,
		server: gin.Default(),
	}
}

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

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Use(CORSMiddleware())

	log.Fatal(router.Run(":" + s.port))
}
