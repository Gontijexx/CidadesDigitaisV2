package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/validation"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		server.DB.SingularTable(true)
		if err != nil {
			fmt.Printf("Cannot connect to database ", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	}

}

func (server *Server) Run() {
	httpServer := &http.Server{
		Addr: config.SERVER_ADDR,

		IdleTimeout:  200 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 300 * time.Millisecond,
	}

	h := server.CreateHandler()

	httpServer.Handler = h

	validation.CreateValidator()
	log.Println("Listening to port 8080")
	log.Fatal("[CONNECTING] ", httpServer.ListenAndServe())

}
