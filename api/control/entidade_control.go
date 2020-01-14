package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/models"
	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/validation"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (server *Server) CreateEntidade(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	entidade := models.Entidade{}

	err = json.Unmarshal(body, &entidade)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(entidade); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	entidadeCreated, err := entidade.SaveEntidade(server.DB)

	if err != nil {

		formattedError := config.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, entidadeCreated.Cnpj))
	responses.JSON(w, http.StatusCreated, entidadeCreated)

}

func (server *Server) GetEntidade() {

}

func (server *Server) GetEntidades() {

}

func (server *Server) UpdateEntidades() {

}

func (server *Server) DeleteEntidades() {

}
