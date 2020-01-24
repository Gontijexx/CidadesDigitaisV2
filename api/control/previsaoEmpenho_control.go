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
	FUNCAO ADICIONAR PREVISAO EMPENHO
=========================  */

func (server *Server) CreatePrevisaoEmpenho(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 18001)

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	//	Estrutura models.PrevisaoEmpenho{} "renomeada"
	previsaoEmpenho := models.PrevisaoEmpenho{}

	//previsaoEmpenho.Prepare()

	//	Unmarshal analisa o JSON recebido e armazena na struct previsaoEmpenho referenciada (&struct)
	err = json.Unmarshal(body, &previsaoEmpenho)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(previsaoEmpenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SavePrevisao_empenho eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	previsaoEmpenhoCreated, err := previsaoEmpenho.SavePrevisaoEmpenho(server.DB)

	//	Retorna um erro caso nao seja possivel salvar entidado no banco de dados
	//	Status 500
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, previsaoEmpenhoCreated.Cod_previsao_empenho))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, previsaoEmpenhoCreated)

}

/*  =========================
	FUNCAO LISTAR PREVISAO EMPENHO POR ID
=========================  */

func (server *Server) GetPrevisaoEmpenhoByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 18002)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	previsaoEmpenhoID armazeza a chave primaria da tabela previsao_empenho
	previsaoEmpenhoID, err := strconv.ParseUint(vars["cod_previsao_empenho"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	previsaoEmpenho := models.PrevisaoEmpenho{}

	//	previsaoEmpenhoGotten recebe o dado buscado no banco de dados
	previsaoEmpenhoGotten, err := previsaoEmpenho.FindPrevisaoEmpenhoByID(server.DB, uint64(previsaoEmpenhoID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, previsaoEmpenhoGotten)

}

/*  =========================
	FUNCAO LISTAR PREVISAO EMPENHO
=========================  */

func (server *Server) GetPrevisaoEmpenho(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 18002)

	previsaoEmpenho := models.PrevisaoEmpenho{}

	//	previsaoEmpenho recebe os dados buscados no banco de dados
	previsaoEmpenhos, err := previsaoEmpenho.FindAllPrevisaoEmpenho(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, previsaoEmpenhos)
}

/*  =========================
	FUNCAO EDITAR PREVISAO EMPENHO
=========================  */

func (server *Server) UpdatePrevisaoEmpenho(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 18003)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	previsaEmepenhoID armazena a chave primaria da tabela entidade
	previsaoEmpenhoID, err := strconv.ParseUint(vars["cod_previsao_empenho"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	previsaoEmpenho := models.PrevisaoEmpenho{}

	//previsaoEmpenho.Prepare()

	err = json.Unmarshal(body, &previsaoEmpenho)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(previsaoEmpenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	// updatePrevisaEmpenho recece a nova previsao_empenho, a que foi alterada
	updatePrevisaoEmpenho, err := previsaoEmpenho.UpdatePrevisaoEmpenho(server.DB, uint64(previsaoEmpenhoID))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updatePrevisaoEmpenho)
}
