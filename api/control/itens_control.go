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
	FUNCAO ADICIONAR ITENS
=========================  */

func (server *Server) CreateItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 23001)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the body, %v\n", err))
		return
	}

	//	Estrutura models.Itens{} "renomeada"
	itens := models.Itens{}

	//	Unmarshal analisa o JSON recebido e armazena na struct itens referenciada (&struct)
	err = json.Unmarshal(body, &itens)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(itens); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveItens eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	itensCreated, err := itens.SaveItens(server.DB)

	/*	Retorna um erro caso nao seja possivel salvar itens no banco de dados
		Status 500	*/
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, itensCreated.CodItem))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, itensCreated)

}

/*  =========================
	FUNCAO LISTAR ITENS POR ID
=========================  */

func (server *Server) GetItensByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 12002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codItem armazena a chave primaria da tabela itens
	codItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codTipoItem armazena a chave primaria da tabela itens
	codTipoItem, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	itens := models.Itens{}

	//	itensGotten recebe o dado buscado no banco de dados
	itensGotten, err := itens.FindItensByID(server.DB, codItem, codTipoItem)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, itensGotten)

}

/*  =========================
	FUNCAO LISTAR TODAS ITENS
=========================  */

func (server *Server) GetAllItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 12002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	itens := models.Itens{}

	//	allItens armazena os dados buscados no banco de dados
	allItens, err := itens.FindAllItens(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allItens)
}

/*  =========================
	FUNCAO EDITAR ITENS
=========================  */

func (server *Server) UpdateItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 12003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codItem armazena a chave primaria da tabela itens
	codItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codTipoItem armazena a chave primaria da tabela itens
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

	itens := models.Itens{}

	err = json.Unmarshal(body, &itens)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(itens); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateItens recebe a nova itens, a que foi alterada
	updateItens, err := itens.UpdateItens(server.DB, codItem, codTipoItem)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateItens)
}

/*  =========================
	FUNCAO DELETAR ITENS
=========================  */

func (server *Server) DeleteItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	err := config.AuthMod(w, r, 12003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	itens := models.Itens{}

	//	codItem armazena a chave primaria da tabela itens
	codItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codTipoItem armazena a chave primaria da tabela itens
	codTipoItem, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = itens.DeleteItens(server.DB, codItem, codTipoItem)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", codItem))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
