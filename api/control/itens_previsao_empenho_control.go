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

func (server *Server) GetItens_previsao_empenhoByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	itens_previsao_empenhoID1, err := strconv.ParseUint(vars["id1"], 10, 64)
	itens_previsao_empenhoID2, err := strconv.ParseUint(vars["id2"], 10, 64)
	itens_previsao_empenhoID3, err := strconv.ParseUint(vars["id3"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	itens_previsao_empenho := models.Itens_previsao_empenho{}

	itens_previsao_empenhoGotten, err := itens_previsao_empenho.FindItens_previsao_empenhoByID(server.DB, uint64(itens_previsao_empenhoID1), uint64(itens_previsao_empenhoID2), uint64(itens_previsao_empenhoID3))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, itens_previsao_empenhoGotten)

}

func (server *Server) GetItens_previsao_empenho(w http.ResponseWriter, r *http.Request) {

	enti := models.Itens_previsao_empenho{}

	itens_previsao_empenho, err := enti.FindAllItens_previsao_empenho(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, itens_previsao_empenho)
}

func (server *Server) UpdateItens_previsao_empenho(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	itens_previsao_empenhoID1, err := strconv.ParseUint(vars["id1"], 10, 64)
	itens_previsao_empenhoID2, err := strconv.ParseUint(vars["id2"], 10, 64)
	itens_previsao_empenhoID3, err := strconv.ParseUint(vars["id3"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itens_previsao_empenho := models.Itens_previsao_empenho{}
	err = json.Unmarshal(body, &itens_previsao_empenho)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(itens_previsao_empenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	updateItens_previsao_empenho, err := itens_previsao_empenho.UpdateItens_previsao_empenho(server.DB, uint64(itens_previsao_empenhoID1), uint64(itens_previsao_empenhoID2), uint64(itens_previsao_empenhoID3))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updateItens_previsao_empenho)
}
