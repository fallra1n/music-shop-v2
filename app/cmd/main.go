package main

import (
	"context"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/handler"
	repository2 "github.com/asssswv/music-shop-v2/app/pkg/repository"
	"github.com/asssswv/music-shop-v2/app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("fail to get password: %s", err.Error())
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
		logrus.Fatalf("fail to init db: %s", err.Error())
	}

	repos := repository2.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(msh.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error ocurred while running http server: %s", err.Error())
		}
	}()

	qiut := make(chan os.Signal, 1)
	signal.Notify(qiut, syscall.SIGTERM, syscall.SIGINT)
	<-qiut

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error ocurred on db connection close: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
