package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Gontijexx/CidadesDigitaisV2/api/config"
	"github.com/Gontijexx/CidadesDigitaisV2/api/models"
	"github.com/Gontijexx/CidadesDigitaisV2/api/responses"
	"github.com/Gontijexx/CidadesDigitaisV2/api/validation"

	"github.com/gorilla/mux"
)

/*  =========================
	FUNCAO ADICIONAR PONTO
=========================  */

func (server *Server) CreatePonto(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13011)
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

	//	Struct's necessarias
	pid := models.Pid{}
	ponto := models.Ponto{}

	//	Unmarshal analisa o JSON recebido e armazena na struct ponto referenciada (&struct)
	err = json.Unmarshal(body, &pid)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	err = json.Unmarshal(body, &ponto)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}
	if err = validation.Validator.Struct(pid); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	if err = validation.Validator.Struct(ponto); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	pidCreated, err := pid.SavePID(server.DB)

	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	ponto.CodPID = pidCreated.CodPid

	//	SavePonto eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	pontoCreated, err := ponto.SavePonto(server.DB)

	//	Retorna um erro caso nao seja possivel salvar ponto no banco de dados
	//	Status 500
	if err != nil {
		_, _ = pid.DeletePID(server.DB, pidCreated.CodPid)
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d/%d/%d", r.Host, r.RequestURI, pontoCreated.CodPonto, pontoCreated.CodCategoria, pontoCreated.CodIbge))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, pontoCreated)

}

/*  =========================
	FUNCAO LISTAR PONTO POR ID
=========================  */

func (server *Server) GetPontoByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13012)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codPonto armazena a chave primaria da tabela ponto
	codPonto, err := strconv.ParseUint(vars["cod_ponto"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codCategoria armazena a chave primaria da tabela ponto
	codCategoria, err := strconv.ParseUint(vars["cod_categoria"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codIbge armazena a chave primaria da tabela ponto
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	ponto := models.Ponto{}

	//	pontoGotten recebe o dado buscado no banco de dados
	pontoGotten, err := ponto.FindPontoByID(server.DB, codPonto, codCategoria, codIbge)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, pontoGotten)

}

/*  =========================
	FUNCAO LISTAR TODAS PONTO
=========================  */

func (server *Server) GetAllPonto(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13012)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	ponto := models.Ponto{}

	//	allPonto armazena os dados buscados no banco de dados
	allPonto, err := ponto.FindAllPonto(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allPonto)
}

/*  =========================
	FUNCAO EDITAR PONTO
=========================  */

func (server *Server) UpdatePonto(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13013)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codPonto armazena a chave primaria da tabela ponto
	codPonto, err := strconv.ParseUint(vars["cod_ponto"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codCategoria armazena a chave primaria da tabela ponto
	codCategoria, err := strconv.ParseUint(vars["cod_categoria"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codIbge armazena a chave primaria da tabela ponto
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	ponto := models.Ponto{}

	err = json.Unmarshal(body, &ponto)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(ponto); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updatePonto recebe a nova ponto, a que foi alterada
	updatePonto, err := ponto.UpdatePonto(server.DB, codPonto, codCategoria, codIbge)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updatePonto)
}

/*  =========================
	FUNCAO DELETAR PONTO
=========================  */

func (server *Server) DeletePonto(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	err := config.AuthMod(w, r, 13013)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	ponto := models.Ponto{}

	//	codPonto armazena a chave primaria da tabela ponto
	codPonto, err := strconv.ParseUint(vars["cod_ponto"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codCategoria armazena a chave primaria da tabela ponto
	codCategoria, err := strconv.ParseUint(vars["cod_categoria"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	codIbge armazena a chave primaria da tabela ponto
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = ponto.DeletePonto(server.DB, codPonto, codCategoria, codIbge)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d/%d/%d", codPonto, codCategoria, codIbge))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
