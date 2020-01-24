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

/*  =========================
	FUNCAO LISTAR ITENS PREVISAO EMPENHO POR ID
=========================  */

/* ACREDITO QUE ITENS PREVISAO EMPENHO PRECISE SER REVISTA NO FUTURO */

func (server *Server) GetItensPrevisaoEmpenhoByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 18002)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	itensPrevisaoEmpenhoID armazena a chave primaria da tabela itens_previsao_empenho
	itensPrevisaoEmpenhoID, err := strconv.ParseUint(vars["cod_previsao_empenho"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	itensPrevisaoEmpenho := models.ItensPrevisaoEmpenho{}

	//	itensPrevisaoEmpenhoGotten recebe o dado buscado no banco de dados
	itensPrevisaoEmpenhoGotten, err := itensPrevisaoEmpenho.FindItensPrevisaoEmpenhoByID(server.DB, uint64(itensPrevisaoEmpenhoID)))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, itensPrevisaoEmpenhoGotten)
}

/*  =========================
	FUNCAO LISTAR ITENS PREVISAO EMPENHO
=========================  */

func (server *Server) GetItensPrevisaoEmpenho(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 18002)

	itensPrevisaoEmpenho := models.ItensPrevisaoEmpenho{}

	//	itensPrevisaoEmpenhos armazena os dados buscados no banco de dados
	itensPrevisaoEmpenhos, err := itensPrevisaoEmpenho.FindAllItensPrevisaoEmpenho(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, itensPrevisaoEmpenhos)
}

/*  =========================
	FUNCAO EDITAR ITENS PREVISAO EMPENHO
=========================  */

func (server *Server) UpdateItensPrevisaoEmpenho(w http.ResponseWriter, r *http.Request) {

	config.AuthMod(w, r, 18003)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	itensPrevisaoEmpenhoID armazena a chave primaria da tabela itens_previsao_empenho
	itensPrevisaoEmpenhoID, err := strconv.ParseUint(vars["cod_previsao_empenho"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itensPrevisaoEmpenho := models.ItensPrevisaoEmpenho{}
	err = json.Unmarshal(body, &itens_previsao_empenho)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(itensPrevisaoEmpenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateEntidade recebe a nova entidade, a que foi alterada
	updateItensPrevisaoEmpenho, err := itensPrevisaoEmpenho.UpdateItensPrevisaoEmpenho(server.DB, uint64(itensPrevisaEempenhoID))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateItensPrevisaoEmpenho)
}
