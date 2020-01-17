package control

import (
	"CidadesDigitaisV2/api/auth"
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/models"
	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/validation"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateReajuste(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	reajuste := models.Reajuste{}

	err = json.Unmarshal(body, &reajuste)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(reajuste); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	reajusteCreated, err := reajuste.SaveReajuste(server.DB)

	if err != nil {

		formattedError := config.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, reajusteCreated.Ano_ref))
	responses.JSON(w, http.StatusCreated, reajusteCreated)

}

func (server *Server) GetReajuste(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	reajusteID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	reajuste := models.Reajuste{}

	reajusteGotten, err := reajuste.FindReajusteByID(server.DB, uint64(reajusteID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, reajusteGotten)

}

func (server *Server) GetReajusteByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	reajuste := models.Reajuste{}
	reajusteGotten, err := reajuste.FindReajusteByID(server.DB, uint64(rId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, reajusteGotten)
}

func (server *Server) UpdateReajustes(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	reajuste := models.Reajuste{}
	err = json.Unmarshal(body, &reajuste)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err = validation.Validator.Struct(reajuste); err != nil {
		log.Printf("[WARN] invalid reajuste information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(rid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	updatedReajuste, err := reajuste.UpdateReajuste(server.DB, uint32(rid))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedReajuste)
}

/*
func (server *Server) DeleteReajuste(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	reajuste := models.Reajuste{}

	rId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(rId) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = reajuste.DeleteReajuste(server.DB, uint32(rId), uint32(rFk))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", rId))
	responses.JSON(w, http.StatusNoContent, "")
}
*/
