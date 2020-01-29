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
	FUNCAO ADICIONAR CD
=========================  */

func (server *Server) CreateCD(w http.ResponseWriter, r *http.Request) {

	//Autorização de Modulo
	err := config.AuthMod(w, r, 13001)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	O metodo RealAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the body, %v\n", err))
	}

	//	Estrutura models.Cd{} "renomeada"
	cd := models.CD{}

	//	Unmarshal analisa o JSON recebido e armazena na struct cd referenciada (&struct)
	err = json.Unmarshal(body, &cd)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(cd); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveCD eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	cdCreated, err := cd.SaveCD(server.DB)

	//	Retorna um erro caso nao seja possivel salvar cd no banco de dados
	//	Status 500
	if err != nil {

		formattedError := config.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, cdCreated.Cod_ibge))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, cdCreated)

}

/*  =========================
	FUNCAO LISTAR CD POR ID
=========================  */

func (server *Server) GetCDByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//Vars retorna a rota das variaveis
	vars := mux.Vars(r)

	//interpreta  a string em uma base de (0, 2 to 36) e tamanho de (0 to 64)
	cdID, err := strconv.ParseUint(vars["cod_ibge"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	cd := models.CD{}

	//vai utilizar o metodo para procurar o resultado de acordo com a chave
	cdGotten, err := cd.FindCDByID(server.DB, uint64(cdID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//retorna um JSON indicando que funcionou corretamente
	responses.JSON(w, http.StatusOK, cdGotten)
}

/*  =========================
	FUNCAO LISTAR TODOS CD
=========================  */

func (server *Server) GetAllCD(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	cd := models.CD{}

	//	cds armazena os dados buscados no banco de dados
	allCD, err := cd.FindAllCD(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allCD)

}

/*  =========================
	FUNCAO EDITAR CD
=========================  */

func (server *Server) UpdateCD(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	cdID armazena a chave primaria da tabela cd
	cdID, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	cd := models.CD{}

	err = json.Unmarshal(body, &cd)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(cd); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateCD recebe a nova cd, a que foi alterada
	updateCD, err := cd.UpdateCD(server.DB, cdID)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateCD)
}
