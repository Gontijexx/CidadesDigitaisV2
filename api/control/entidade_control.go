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
	FUNCAO ADICIONAR ENTIDADE
=========================  */

func (server *Server) CreateEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 12001)

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	//	Estrutura models.Entidade{} "renomeada"
	entidade := models.Entidade{}

	//	Unmarshal analisa o JSON recebido e armazena na struct entidade referenciada (&struct)
	err = json.Unmarshal(body, &entidade)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
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

		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[Error] We couldn't save Entidade, Check server details"))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, entidadeCreated.Cnpj))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, entidadeCreated)

}

/*  =========================
	FUNCAO LISTAR ENTIDADE POR ID
=========================  */

func (server *Server) GetEntidadeByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 12002)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	entidadeID armazena a chave primaria da tabela entidade
	entidadeID, err := strconv.ParseUint(vars["cnpj"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	entidade := models.Entidade{}

	//	entidadeGotten recebe o dado buscado no banco de dados
	entidadeGotten, err := entidade.FindEntidadeByID(server.DB, uint64(entidadeID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, entidadeGotten)

}

/*  =========================
	FUNCAO LISTAR TODAS ENTIDADE
=========================  */

func (server *Server) GetEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 12002)

	entidade := models.Entidade{}

	//	entidades armazena os dados buscados no banco de dados
	entidades, err := entidade.FindEntidades(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, entidades)
}

/*  =========================
	FUNCAO EDITAR ENTIDADE
=========================  */

func (server *Server) UpdateEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 12003)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	entidadeID armazena a chave primaria da tabela entidade
	entidadeID, err := strconv.ParseUint(vars["cnpj"], 10, 64)
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

	//	updateEntidade recebe a nova entidade, a que foi alterada
	updateEntidade, err := entidade.UpdateEntidade(server.DB, entidadeID)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateEntidade)
}

/*  =========================
	FUNCAO DELETAR ENTIDADE
=========================  */

func (server *Server) DeleteEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	config.AuthMod(w, r, 12003)

	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	entidade := models.Entidade{}

	//	entidadeID armazena a chave primaria da tabela entidade
	entidadeID, err := strconv.ParseUint(vars["cnpj"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = entidade.DeleteEntidade(server.DB, uint64(entidadeID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", entidadeID))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
