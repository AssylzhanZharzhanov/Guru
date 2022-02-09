package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/zharzhanov/myguru/app"
	"gitlab.com/zharzhanov/myguru/db/postgres"
	"gitlab.com/zharzhanov/myguru/pkg/handler"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
	"gitlab.com/zharzhanov/myguru/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	srv := new(app.Server)

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName: viper.GetString("database.name"),
		SSLMode: viper.GetString("database.ssl_mode"),
	})

	if err != nil {
		log.Fatalf("Error in starting database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := viper.GetString("server.port")
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			log.Fatalf("Error in starting server: %s", err.Error())
		}
	}()
	logrus.Print("Server started at " + port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
