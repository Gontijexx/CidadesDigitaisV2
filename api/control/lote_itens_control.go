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

func (server *Server) GetLote_itensByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	lote_itensID1, err := strconv.ParseUint(vars["id1"], 10, 64)
	lote_itensID2, err := strconv.ParseUint(vars["id2"], 10, 64)
	lote_itensID3, err := strconv.ParseUint(vars["id3"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	lote_itens := models.Lote_itens{}

	lote_itensGotten, err := lote_itens.FindLote_itensByID(server.DB, uint64(lote_itensID1), uint64(lote_itensID2), uint64(lote_itensID3))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, lote_itensGotten)

}

func (server *Server) GetLote_itens(w http.ResponseWriter, r *http.Request) {

	enti := models.Lote_itens{}

	lote_itens, err := enti.FindAllLote_itens(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, lote_itens)
}

func (server *Server) UpdateLote_itens(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lote_itensID1, err := strconv.ParseUint(vars["id1"], 10, 64)
	lote_itensID2, err := strconv.ParseUint(vars["id2"], 10, 64)
	lote_itensID3, err := strconv.ParseUint(vars["id3"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	lote_itens := models.Lote_itens{}
	err = json.Unmarshal(body, &lote_itens)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(lote_itens); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	updateLote_itens, err := lote_itens.UpdateLote_itens(server.DB, uint64(lote_itensID1), uint64(lote_itensID2), uint64(lote_itensID3))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updateLote_itens)
}
