package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/models"
	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/validation"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) GetCd_itensByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	cd_itensID1, err := strconv.ParseUint(vars["id1"], 10, 64)
	cd_itensID2, err := strconv.ParseUint(vars["id2"], 10, 64)
	cd_itensID3, err := strconv.ParseUint(vars["id3"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	cd_itens := models.Cd_itens{}

	cd_itensGotten, err := cd_itens.FindCd_itensByID(server.DB, uint64(cd_itensID1), uint64(cd_itensID2), uint64(cd_itensID3))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, cd_itensGotten)

}

func (server *Server) GetCd_itens(w http.ResponseWriter, r *http.Request) {

	enti := models.Cd_itens{}

	cd_itens, err := enti.FindAllCd_itens(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, cd_itens)
}

func (server *Server) UpdateCd_itens(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cd_itensID1, err := strconv.ParseUint(vars["id1"], 10, 64)
	cd_itensID2, err := strconv.ParseUint(vars["id2"], 10, 64)
	cd_itensID3, err := strconv.ParseUint(vars["id3"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cd_itens := models.Cd_itens{}
	err = json.Unmarshal(body, &cd_itens)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(cd_itens); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	updateCd_itens, err := cd_itens.UpdateCd_itens(server.DB, uint64(cd_itensID1), uint64(cd_itensID2), uint64(cd_itensID3))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updateCd_itens)
}
