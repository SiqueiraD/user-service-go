package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/siqueirad/user-service-go/routes"
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

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Fatal(router.Run(":" + s.port))
}
