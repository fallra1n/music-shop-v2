package music_shop

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) InitRouter(router *gin.Engine) {
	s.router = router
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

func (s *Server) RunServer(port string) {
	_ = s.router.Run("localhost:" + port)
}

func NewServer() *Server {
	return new(Server)
}

var Srv = NewServer()
