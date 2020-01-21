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

//	Funcao criar uma entidade no banco de dados
func (server *Server) CreateEntidade(w http.ResponseWriter, r *http.Request) {

	//	O metodo RealAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	//	Estrutura models.Entidade{} "renomeada"
	entidade := models.Entidade{}

	//	Unmarshal analisa o JSON recebido e armazena na struct entidade referenciada (&struct)
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

	//	SaveEntidade eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	entidadeCreated, err := entidade.SaveEntidade(server.DB)

	//	Retorna um erro caso nao seja possivel salvar entidado no banco de dados
	//	Status 500
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, entidadeCreated.Cnpj))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, entidadeCreated)

}

func (server *Server) GetEntidadeByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	entidadeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	entidade := models.Entidade{}

	entidadeGotten, err := entidade.FindEntidadeByID(server.DB, uint64(entidadeID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, entidadeGotten)

}

func (server *Server) GetEntidades(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) UpdateEntidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	entidadeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
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

	updateEntidade, err := entidade.UpdateEntidade(server.DB, uint64(entidadeID))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updateEntidade)
}

func (server *Server) DeleteEntidade(w http.ResponseWriter, r *http.Request) {

	// vars recebe o ID contido na URL
	vars := mux.Vars(r)

	entidade := models.Entidade{}

	entidadeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = entidade.DeleteEntidade(server.DB, uint64(entidadeID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", entidadeID))
	responses.JSON(w, http.StatusNoContent, "")
}
