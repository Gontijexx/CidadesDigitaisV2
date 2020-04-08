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
	FUNCAO ADICIONAR OTB FATURA
=========================  */

func (server *Server) CreateOTBFatura(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 16011)
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

	//	Estrutura models.FaturaOTB{} "renomeada"
	faturaOTB := models.FaturaOTB{}

	//	Unmarshal analisa o JSON recebido e armazena na struct faturaOTB referenciada (&struct)
	err = json.Unmarshal(body, &faturaOTB)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(faturaOTB); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveFaturaOTB eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	faturaOTBCreated, err := faturaOTB.SaveOTBFatura(server.DB)

	/*	Retorna um erro caso nao seja possivel salvar faturaOTB no banco de dados
		Status 500	*/
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, faturaOTBCreated.CodOtb))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, faturaOTBCreated)

}

/*  =========================
	FUNCAO LISTAR TODAS OTB FATURA
=========================  */

func (server *Server) GetOTBFatura(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 16012)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codOTB armazena a chave primaria da tabela entidade
	codOTB, err := strconv.ParseUint(vars["cod_otb"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	faturaOTB := models.FaturaOTB{}

	//	allFaturaOTB armazena os dados buscados no banco de dados
	allFaturaOTB, err := faturaOTB.FindOTBFatura(server.DB, codOTB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allFaturaOTB)
}
