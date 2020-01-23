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

/*  =========================
	FUNCAO ADICIONAR LOTE
=========================  */

func (server *Server) CreateLote(w http.ResponseWriter, r *http.Request) {

	//Autorização de Modulo
	config.AuthMod(w, r, 14001)

	//	O metodo RealAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	//	Estrutura models.Lote{} "renomeada"
	lote := models.Lote{}

	//	Unmarshal analisa o JSON recebido e armazena na struct lote referenciada (&struct)
	err = json.Unmarshal(body, &lote)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(lote); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveLote eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	loteCreated, err := lote.SaveLote(server.DB)

	//	Retorna um erro caso nao seja possivel salvar lote no banco de dados
	//	Status 500
	if err != nil {

		formattedError := config.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, loteCreated.Cnpj))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, loteCreated)

}

/*  =========================
	FUNCAO LISTAR LOTE
=========================  */

func (server *Server) GetLote(w http.ResponseWriter, r *http.Request) {

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	loteID armazena a chave primaria da tabela entidade
	loteID, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	lote := models.Lote{}

	//	loteGotten recebe o dado buscado no banco de dados
	loteGotten, err := lote.FindLoteByID(server.DB, uint64(loteID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, loteGotten)

}

/*  =========================
	FUNCAO LISTAR LOTE POR ID
=========================  */

func (server *Server) GetLoteByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lId, err := strconv.ParseUint(vars["cod_lote"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	lote := models.Lote{}
	loteGotten, err := lote.FindLoteByID(server.DB, uint64(lId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, loteGotten)
}

/*  =========================
	FUNCAO ATUALIZAR LOTE
=========================  */

func (server *Server) UpdateLote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	lote := models.Lote{}
	err = json.Unmarshal(body, &lote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err = validation.Validator.Struct(lote); err != nil {
		log.Printf("[WARN] invalid lote information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(lid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	updatedLote, err := lote.UpdateLote(server.DB, uint64(lid))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedLote)
}
