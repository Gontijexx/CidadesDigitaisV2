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
	FUNCAO ADICIONAR NATUREZA_DESPESA
=========================  */

func (server *Server) CreateNaturezaDespesa(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 25001)
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

	//	Estrutura models.NaturezaDespesa{} "renomeada"
	naturezaDespesa := models.NaturezaDespesa{}

	//	Unmarshal analisa o JSON recebido e armazena na struct natureza_despesa referenciada (&struct)
	err = json.Unmarshal(body, &naturezaDespesa)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(naturezaDespesa); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveNaturezaDespesa eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	naturezaDespesaCreated, err := naturezaDespesa.SaveNaturezaDespesa(server.DB)

	/*	Retorna um erro caso nao seja possivel salvar natureza_despesa no banco de dados
		Status 500	*/
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, naturezaDespesaCreated.CodNaturezaDespesa))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, naturezaDespesaCreated)

}

/*  =========================
	FUNCAO LISTAR NATUREZA_DESPESA POR ID
=========================  */

func (server *Server) GetNaturezaDespesaByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 25002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codNaturezaDespesa armazena a chave primaria da tabela natureza_despesa
	codNaturezaDespesa, err := strconv.ParseUint(vars["cod_natureza_despesa"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	naturezaDespesa := models.NaturezaDespesa{}

	//	naturezaDespesaGotten recebe o dado buscado no banco de dados
	naturezaDespesaGotten, err := naturezaDespesa.FindNaturezaDespesaByID(server.DB, codNaturezaDespesa)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, naturezaDespesaGotten)

}

/*  =========================
	FUNCAO LISTAR TODAS NATUREZA_DESPESA
=========================  */

func (server *Server) GetAllNaturezaDespesa(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 25002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	naturezaDespesa := models.NaturezaDespesa{}

	//	allNaturezaDespesa armazena os dados buscados no banco de dados
	allNaturezaDespesa, err := naturezaDespesa.FindAllNaturezaDespesa(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allNaturezaDespesa)
}

/*  =========================
	FUNCAO EDITAR NATUREZA_DESPESA
=========================  */

func (server *Server) UpdateNaturezaDespesa(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 25003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codNaturezaDespesa armazena a chave primaria da tabela natureza_despesa
	codNaturezaDespesa, err := strconv.ParseUint(vars["cod_natureza_despesa"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	naturezaDespesa := models.NaturezaDespesa{}

	err = json.Unmarshal(body, &naturezaDespesa)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(naturezaDespesa); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateNaturezaDespesa recebe a nova natureza_despesa, a que foi alterada
	updateNaturezaDespesa, err := naturezaDespesa.UpdateNaturezaDespesa(server.DB, codNaturezaDespesa)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateNaturezaDespesa)
}

/*  =========================
	FUNCAO DELETAR NATUREZA_DESPESA
=========================  */

func (server *Server) DeleteNaturezaDespesa(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	err := config.AuthMod(w, r, 25003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	naturezaDespesa := models.NaturezaDespesa{}

	//	codNaturezaDespesa armazena a chave primaria da tabela natureza_despesa
	codNaturezaDespesa, err := strconv.ParseUint(vars["cod_natureza_despesa"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = naturezaDespesa.DeleteNaturezaDespesa(server.DB, codNaturezaDespesa)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", codNaturezaDespesa))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
