package server

import (
	"log"

	"github.com/douglasandradeee/web_api_gin/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine // Inicia o servidor Gin
}

func NewServer() Server {
	return Server{
		port:   "9000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("Server is running on port: ", s.port)
	log.Fatal(router.Run(":" + s.port))

}
