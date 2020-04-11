package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"

	"CidadesDigitaisV2/api/control"

	"github.com/joho/godotenv"
)

/*	=========================
		COMENTAR
=========================	*/

var server = control.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Router = mux.NewRouter()

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run()

}
