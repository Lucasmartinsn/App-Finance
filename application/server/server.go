package server

import (
	"development/application/fiance/server/router"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// A struct Server armazena os dados que vai iniciar o servidor
type Server struct {
	port   string
	server *gin.Engine
	origin string
}

// Essa função vai carregar os dados que vai ser usado para inicializar o Servidor
func NewServer() Server {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Falhar ao carregar variaveis de ambiente")
	}
	return Server{
		port:   os.Getenv("API_PORT"),
		server: gin.Default(),
		origin: os.Getenv("API_LOCALHOST"),
	}
}

// Vai criar o servidor Gin
func (s *Server) Run() {
	routers := router.ConfigRouter(s.server)
	log.Print("Server running on ", s.port)
	log.Fatal(routers.Run(fmt.Sprintf("%s:%s", s.origin, s.port)))
}
