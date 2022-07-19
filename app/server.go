package app

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 MB
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

//func (s *Server) GetRouter() *gin.Engine {
//	return s.router
//}
//
//func (s *Server) RunServer(port string) {
//	_ = s.router.Run("localhost:" + port)
//}
//
//func NewServer() *Server {
//	return new(Server)
//}
//
//var Srv = NewServer()
