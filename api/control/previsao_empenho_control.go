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

func (server *Server) CreatePrevisao_empenho(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)

	//Autorização de Modulo
	config.AuthMod(w, r, 18001)

	//	O metodo RealAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	//	Estrutura models.Previsao_empenho{} "renomeada"
	previsao_empenho := models.Previsao_empenho{}

	//	Unmarshal analisa o JSON recebido e armazena na struct previsao_empenho referenciada (&struct)
	err = json.Unmarshal(body, &previsao_empenho)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	log.Printf("%v", previsao_empenho)

	if err = validation.Validator.Struct(previsao_empenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SavePrevisao_empenho eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	previsao_empenhoCreated, err := previsao_empenho.SavePrevisao_empenho(server.DB)

	//	Retorna um erro caso nao seja possivel salvar entidado no banco de dados
	//	Status 500
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, previsao_empenhoCreated.Cod_previsao_empenho))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, previsao_empenhoCreated)

}

func (server *Server) GetPrevisao_empenhoByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	previsao_empenhoID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	previsao_empenho := models.Previsao_empenho{}

	previsao_empenhoGotten, err := previsao_empenho.FindPrevisao_empenhoByID(server.DB, uint64(previsao_empenhoID))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, previsao_empenhoGotten)

}

func (server *Server) GetPrevisao_empenhos(w http.ResponseWriter, r *http.Request) {

	enti := models.Previsao_empenho{}

	previsao_empenho, err := enti.FindAllPrevisao_empenho(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, previsao_empenho)
}

func (server *Server) UpdatePrevisao_empenho(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	previsao_empenhoID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	previsao_empenho := models.Previsao_empenho{}
	err = json.Unmarshal(body, &previsao_empenho)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(previsao_empenho); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	updatePrevisao_empenho, err := previsao_empenho.UpdatePrevisao_empenho(server.DB, uint64(previsao_empenhoID))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatePrevisao_empenho)
}
