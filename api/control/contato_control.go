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
	FUNCAO ADICIONAR CONTATO
=========================  */

func (server *Server) AddContato(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	//config.AuthMod(w, r, 12001)

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	//	Estrutura models.Contato{} "renomeada"
	contato := models.Contato{}

	/*	O metodo Prepare deve ser chamado em metodos de POST e PUT
		a fim de preparar os dados a serem recebidos pelo banco de dados	*/
	contato.Prepare()

	//	Unmarshal analisa o JSON recebido e armazena na struct contato referenciada (&struct)
	err = json.Unmarshal(body, &contato)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(contato); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveContato eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	contatoCreated, err := contato.SaveContato(server.DB)

	//	Retorna um erro caso nao seja possivel salvar entidado no banco de dados
	//	Status 500
	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[Error] We couldn't save Contato, Check server details"))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, contatoCreated.Cod_contato))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, contatoCreated)

}

/*  =========================
	FUNCAO LISTAR CONTATO POR ID
=========================  */

func (server *Server) GetContato(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	//config.AuthMod(w, r, 12002)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	contatoID armazena a chave primaria da tabela contato
	contatoID, err := strconv.ParseUint(vars["cod_contato"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	contato := models.Contato{}

	//	contatoGotten recebe o dado buscado no banco de dados
	contatoGotten, err := contato.FindContato(server.DB, uint64(contatoID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, contatoGotten)

}

/*  =========================
	FUNCAO LISTAR CONTATOS
=========================  */

func (server *Server) GetContatos(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	//config.AuthMod(w, r, 12002)

	contato := models.Contato{}

	//	contatos armazena os dados buscados no banco de dados
	contatos, err := contato.FindContatos(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, contatos)
}

/*  =========================
	FUNCAO EDITAR CONTATO
=========================  */

func (server *Server) UpdateContato(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	//config.AuthMod(w, r, 12003)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	contatoID armazena a chave primaria da tabela contato
	contatoID, err := strconv.ParseUint(vars["cdo_contato"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	contato := models.Contato{}

	contato.Prepare()

	err = json.Unmarshal(body, &contato)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(contato); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateContato recebe a nova contato, a que foi alterada
	updateContato, err := contato.UpdateContato(server.DB, contatoID)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateContato)
}

/*  =========================
	FUNCAO DELETAR CONTATO
=========================  */

func (server *Server) DeleteContato(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	//config.AuthMod(w, r, 12003)

	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	contato := models.Contato{}

	//	contatoID armazena a chave primaria da tabela contato
	contatoID, err := strconv.ParseUint(vars["cod_contato"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = contato.DeleteContato(server.DB, uint64(contatoID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", contatoID))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
