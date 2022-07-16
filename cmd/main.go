package main

import (
	msh "github.com/asssswv/music-shop-v2"
	"github.com/asssswv/music-shop-v2/pkg/handler"
	"github.com/asssswv/music-shop-v2/pkg/repository"
	"github.com/asssswv/music-shop-v2/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	msh.Srv.InitRouter(handlers.InitRoutes())
	msh.Srv.RunServer()
}
