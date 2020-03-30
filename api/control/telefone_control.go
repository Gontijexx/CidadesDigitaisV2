package control

import (
	"encoding/json"
	"fmt"
	"github.com/Gontijexx/CidadesDigitaisV2/api/config"
	"github.com/Gontijexx/CidadesDigitaisV2/api/models"
	"github.com/Gontijexx/CidadesDigitaisV2/api/responses"
	"github.com/Gontijexx/CidadesDigitaisV2/api/validation"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*  =========================
	FUNCAO ADICIONAR TELEFONE
=========================  */

func (server *Server) CreateTelefone(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 12001)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the body, %v\n", err))
	}

	//	Estrutura models.Telefone{} "renomeada"
	telefone := models.Telefone{}

	/*	O metodo Prepare deve ser chamado em metodos de POST e PUT
		a fim de preparar os dados a serem recebidos pelo banco de dados	*/
	//telefone.Prepare()

	//	Unmarshal analisa o JSON recebido e armazena na struct telefone referenciada (&struct)
	err = json.Unmarshal(body, &telefone)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(telefone); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveTelefone eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	telefoneCreated, err := telefone.SaveTelefone(server.DB)

	//	Retorna um erro caso nao seja possivel salvar entidado no banco de dados
	//	Status 500
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, telefoneCreated.CodTelefone))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, telefoneCreated)

}

/*  =========================
	FUNCAO LISTAR TODOS TELEFONE
=========================  */

func (server *Server) GetAllTelefone(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 12002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	telefone := models.Telefone{}

	//	telefones armazena os dados buscados no banco de dados
	allTelefone, err := telefone.FindAllTelefone(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allTelefone)
}

/*  =========================
	FUNCAO DELETAR TELEFONE
=========================  */

func (server *Server) DeleteTelefone(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	err := config.AuthMod(w, r, 12003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	telefone := models.Telefone{}

	//	codTelefone armazena a chave primaria da tabela telefone
	codTelefone, err := strconv.ParseUint(vars["cod_telefone"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = telefone.DeleteTelefone(server.DB, codTelefone)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", codTelefone))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
