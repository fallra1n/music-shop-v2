package main

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/handler"
	repository2 "github.com/asssswv/music-shop-v2/app/pkg/repository"
	"github.com/asssswv/music-shop-v2/app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("fail to get password: %s", err.Error())
	}

	db, err := repository2.NewPostgresDB(repository2.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("fail to init db: %s", err.Error())
	}

	repos := repository2.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	msh.Srv.InitRouter(handlers.InitRoutes())
	msh.Srv.RunServer(viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
