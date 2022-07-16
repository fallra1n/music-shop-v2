package music_shop

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
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

func (s *Server) RunServer() {
	if err := InitConfig(); err != nil {
		log.Fatalf("ooops: %s", err.Error())
	}
	_ = s.router.Run("localhost:" + viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func NewServer() *Server {
	return new(Server)
}

var Srv = NewServer()
