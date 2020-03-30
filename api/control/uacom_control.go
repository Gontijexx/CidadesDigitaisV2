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

//time.Now().Format(time.RFC3339)
/*  =========================
	FUNCAO ADICIONAR UACOM
=========================  */

func (server *Server) CreateUacom(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13101)
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

	//	Estrutura models.Uacom{} "renomeada"
	uacom := models.Uacom{}

	//	Unmarshal analisa o JSON recebido e armazena na struct uacom referenciada (&struct)
	err = json.Unmarshal(body, &uacom)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(uacom); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveUacom eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	uacomCreated, err := uacom.SaveUacom(server.DB)

	/*	Retorna um erro caso nao seja possivel salvar uacom no banco de dados
		Status 500	*/
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d/%d", r.Host, r.RequestURI, uacomCreated.CodIbge, uacomCreated.Data))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, uacomCreated)

}

/*  =========================
	FUNCAO LISTAR UACOM POR ID
=========================  */

func (server *Server) GetUacomByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13102)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codIbge armazena a chave primaria da tabela uacom
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	data armazena a chave primaria da tabela uacom
	data := vars["data"]
	fmt.Print(data)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	uacom := models.Uacom{}

	//	uacomGotten recebe o dado buscado no banco de dados
	uacomGotten, err := uacom.FindUacomByID(server.DB, codIbge, data)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, uacomGotten)

}

/*  =========================
	FUNCAO LISTAR TODAS UACOM
=========================  */

func (server *Server) GetAllUacom(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13102)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	uacom := models.Uacom{}

	//	allUacom armazena os dados buscados no banco de dados
	allUacom, err := uacom.FindAllUacom(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allUacom)
}

/*  =========================
	FUNCAO EDITAR UACOM
=========================  */

func (server *Server) UpdateUacom(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13103)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codIbge armazena a chave primaria da tabela uacom
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	data armazena a chave primaria da tabela uacom
	data := vars["data"]
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	uacom := models.Uacom{}

	err = json.Unmarshal(body, &uacom)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(uacom); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateUacom recebe a nova uacom, a que foi alterada
	updateUacom, err := uacom.UpdateUacom(server.DB, codIbge, data)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateUacom)
}

/*  =========================
	FUNCAO DELETAR UACOM
=========================  */

func (server *Server) DeleteUacom(w http.ResponseWriter, r *http.Request) {

	err := config.AuthMod(w, r, 13103)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	uacom := models.Uacom{}

	//	codIbge armazena a chave primaria da tabela uacom
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	data armazena a chave primaria da tabela uacom
	data := vars["data"]
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}
	_, err = uacom.DeleteUacom(server.DB, codIbge, data)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d/%d", codIbge, data))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
