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
	FUNCAO LISTAR ITENS PREVISAO EMPENHO POR ID
=========================  */

func (server *Server) GetItensPrevisaoEmpenhoByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 18002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codPrevisaoEmpenho armazena a chave primaria da tabela itens_previsao_empenho
	codPrevisaoEmpenho, err := strconv.ParseUint(vars["cod_previsao_empenho"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}
	//	codItem armazena a chave primaria da tabela itens_previsao_empenho
	codItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codTipoItem armazena a chave primaria da tabela itens_previsao_empenho
	codTipoItem, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	itensPrevisaoEmpenho := models.ItensPrevisaoEmpenho{}

	//	itensPrevisaoEmpenhoGotten recebe o dado buscado no banco de dados
	itensPrevisaoEmpenhoGotten, err := itensPrevisaoEmpenho.FindItensPrevisaoEmpenhoByID(server.DB, codPrevisaoEmpenho, codItem, codTipoItem)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, itensPrevisaoEmpenhoGotten)
}

/*  =========================
	FUNCAO EDITAR ITENS PREVISAO EMPENHO
=========================  */

func (server *Server) UpdateItensPrevisaoEmpenho(w http.ResponseWriter, r *http.Request) {

	err := config.AuthMod(w, r, 18003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codPrevisaoEmpenho armazena a chave primaria da tabela itens_previsao_empenho
	codPrevisaoEmpenho, err := strconv.ParseUint(vars["cod_previsao_empenho"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}
	//	codItem armazena a chave primaria da tabela itens_previsao_empenho
	codItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codTipoItem armazena a chave primaria da tabela itens_previsao_empenho
	codTipoItem, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	itensPrevisaoEmpenho := models.ItensPrevisaoEmpenho{}
	err = json.Unmarshal(body, &itensPrevisaoEmpenho)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(itensPrevisaoEmpenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateItensPrevisaoEmpenho recebe a nova itens previsao empenho, a que foi alterada
	updateItensPrevisaoEmpenho, err := itensPrevisaoEmpenho.UpdateItensPrevisaoEmpenho(server.DB, codPrevisaoEmpenho, codItem, codTipoItem)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateItensPrevisaoEmpenho)
}
