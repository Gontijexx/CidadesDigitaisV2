package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Gontijexx/CidadesDigitaisV2/api/config"
	"github.com/Gontijexx/CidadesDigitaisV2/api/models"
	"github.com/Gontijexx/CidadesDigitaisV2/api/responses"
	"github.com/Gontijexx/CidadesDigitaisV2/api/validation"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*  =========================
	FUNCAO ADICIONAR ENTIDADE
=========================  */

func (server *Server) CreateEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	if err := config.AuthMod(w, r, 12001); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	entidade := models.Entidade{}
	logEntidade := models.Log{}

	//	Extrai o cod_usuario do body
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the body, %v\n", err))
		return
	}

	//	Unmarshal analisa o JSON recebido e armazena na struct entidade referenciada (&struct)
	err = json.Unmarshal(body, &entidade)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	//	Validacao da struct recebida
	if err = validation.Validator.Struct(entidade); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	Validacao de CNPJ
	if result := validation.ValidationCNPJ(entidade.Cnpj); result != true {
		log.Printf("[FATAL] invalid CNPJ!")
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	logEntidade.ConcatEntidade(server.DB, entidade.Cnpj)

	//	SaveEntidade eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	entidadeCreated, err := entidade.SaveEntidade(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	logEntidade.LogEntidade(server.DB, entidade.Cnpj, "entidade", "i", tokenID)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, entidadeCreated.Cnpj))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, entidadeCreated)

}

/*  =========================
	FUNCAO LISTAR ENTIDADE POR ID
=========================  */

func (server *Server) GetEntidadeByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	if err := config.AuthMod(w, r, 12002); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	cnpj armazena a chave primaria da tabela entidade
	cnpj := vars["cnpj"]

	entidade := models.Entidade{}

	//	entidadeGotten recebe o dado buscado no banco de dados
	entidadeGotten, err := entidade.FindEntidadeByID(server.DB, cnpj)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, entidadeGotten)

}

/*  =========================
	FUNCAO LISTAR TODAS ENTIDADE
=========================  */

func (server *Server) GetAllEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	if err := config.AuthMod(w, r, 12002); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	entidade := models.Entidade{}

	//	allEntidade armazena os dados buscados no banco de dados
	allEntidade, err := entidade.FindAllEntidade(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allEntidade)
}

/*  =========================
	FUNCAO EDITAR ENTIDADE
=========================  */

func (server *Server) UpdateEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	if err := config.AuthMod(w, r, 12003); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	cnpj armazena a chave primaria da tabela entidade
	cnpj := vars["cnpj"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	//	Extrai o cod_usuario do body
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	entidade := models.Entidade{}
	logEntidade := models.Log{}

	err = json.Unmarshal(body, &entidade)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(entidade); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	logEntidade.ConcatEntidade(server.DB, cnpj)

	//	updateEntidade recebe a nova entidade, a que foi alterada
	updateEntidade, err := entidade.UpdateEntidade(server.DB, cnpj)

	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	logEntidade.LogEntidade(server.DB, cnpj, "entidade", "u", tokenID)

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateEntidade)
}

/*  =========================
	FUNCAO DELETAR ENTIDADE
=========================  */

func (server *Server) DeleteEntidade(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	if err := config.AuthMod(w, r, 12003); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	entidade := models.Entidade{}
	logEntidade := models.Log{}

	//	Extrai o cod_usuario do body
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//	cnpj armazena a chave primaria da tabela entidade
	cnpj := vars["cnpj"]

	logEntidade.ConcatEntidade(server.DB, cnpj)

	// 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	//	Caso nao seja possivel deletar o dado especificado tratamos o erro

	_, err = entidade.DeleteEntidade(server.DB, cnpj)

	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	logEntidade.LogEntidade(server.DB, cnpj, "entidade", "d", tokenID)

	w.Header().Set("Entity", fmt.Sprintf("%d", cnpj))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}

/*  =========================
	FUNCAO LISTAR ENTIDADE.CNPJ E ENTIDADE.NOME
=========================  */

func (server *Server) GetEntidadeIDandName(w http.ResponseWriter, r *http.Request) {

	entidade := models.Entidade{}

	//	entidadeGotten recebe o dado buscado no banco de dados
	entidadeGotten, err := entidade.FindEntidadeIDAndName(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	bytes, _ := json.Marshal(entidadeGotten)

	w.Write(bytes)
}
