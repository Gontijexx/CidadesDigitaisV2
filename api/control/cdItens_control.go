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
	"strconv"

	"github.com/gorilla/mux"
)

/*  =========================
	FUNCAO LISTAR CD_ITENS POR ID
=========================  */

func (server *Server) GetCdItem(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 13022)

	//Vars retorna a rota das variaveis
	vars := mux.Vars(r)

	//interpreta  a string em uma base de (0, 2 to 36) e tamanho de (0 to 64)
	cd_itensID1, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	cd_itensID2, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	cd_itensID3, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	cd_itens := models.Cd_itens{}

	//vai utilizar o metodo para procurar o resultado de acordo com a chave
	cd_itensGotten, err := cd_itens.FindCdItem(server.DB, uint64(cd_itensID1), uint64(cd_itensID2), uint64(cd_itensID3))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//retorna um JSON indicando que funcionou corretamente
	responses.JSON(w, http.StatusOK, cd_itensGotten)

}

/*  =========================
	FUNCAO LISTAR CD_ITENS
=========================  */

func (server *Server) GetCdItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 13022)

	cd_item := models.Cd_itens{}

	//	cd_itens armazena os dados buscados no banco de dados
	cd_itens, err := cd_item.FindCdItens(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, cd_itens)
}

/*  =========================
	FUNCAO ATUALIZAR CD_ITENS
=========================  */

func (server *Server) UpdateCdItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 13023)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	cd_itensIDs armazena a chave primaria da tabela cd_itens
	cd_itensID1, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	cd_itensID2, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	cd_itensID3, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)
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

	//cd_itens.Prepare()

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

	//	UpdateCdItem recebe a nova cd_itens, a que foi alterada
	UpdateCdItem, err := cd_itens.UpdateCdItem(server.DB, uint64(cd_itensID1), uint64(cd_itensID2), uint64(cd_itensID3))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, UpdateCdItem)
}
