package api

import (
	"github.com/spf13/viper"
	"gitlab.com/zharzhanov/myguru/app"
	"log"
)

func main(){

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}

	srv := new(app.Server)

	port := viper.GetString("server.port")
	log.Println(port)
	if err := srv.Run(port); err != nil {
		log.Fatalf("Error in starting server: %s", err.Error())
	}
}
