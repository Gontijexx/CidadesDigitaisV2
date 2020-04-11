package control

import (
	"html/template"
	"net/http"

	"CidadesDigitaisV2/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
	homepageTpl := template.Must(template.New("homepage_view").ParseFiles("../Front/home.html"))
	homepageTpl.Execute(w, nil)

	return
}
