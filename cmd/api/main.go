package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/zharzhanov/myguru/app"
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

	port := viper.GetString("server.port")
	go func() {
		if err := srv.Run(port); err != nil {
			log.Fatalf("Error in starting server: %s", err.Error())
		}
	}()
	logrus.Print("Server started at " + port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
