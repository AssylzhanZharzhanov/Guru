package api

import (
	"gitlab.com/zharzhanov/myguru/app"
	"log"
)

func main(){

	srv := new(app.App)

	//port := os.Getenv("port")
	port := "8000"
	if err := srv.Run(port); err != nil {
		log.Fatalf("Error in starting server: %s", err.Error())
	}
}
