package control


import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	

	"CidadesDigitaisV2/api/models"

	"CidadesDigitaisV2/api/validation"
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
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	

}

func (server *Server) Run() {
	httpServer = &http.Server{
		Addr: config.SERVER_ADDR,

		IdleTimeout:  200 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	}

	h := CreateHandler()

	httpServer.Handler = h

	validation.CreateValidator()
	log.Println("Listening to port 8080")
	log.Fatal(httpServer.ListenAndServe())

}