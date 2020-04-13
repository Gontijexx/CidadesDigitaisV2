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
)

/*  =========================
	FUNCAO ADICIONAR UACOM_ASSUNTO
=========================  */

func (server *Server) CreateUacomAssunto(w http.ResponseWriter, r *http.Request) {

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

	//	Estrutura models.UacomAssunto{} "renomeada"
	uacomAssunto := models.UacomAssunto{}

	//	Unmarshal analisa o JSON recebido e armazena na struct uacomAssunto referenciada (&struct)
	err = json.Unmarshal(body, &uacomAssunto)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(uacomAssunto); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveUacomAssunto eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	uacomAssuntoCreated, err := uacomAssunto.SaveUacomAssunto(server.DB)

	/*	Retorna um erro caso nao seja possivel salvar uacomAssunto no banco de dados
		Status 500	*/
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d/%d/%d", r.Host, r.RequestURI, uacomAssuntoCreated.CodIbge, uacomAssuntoCreated.Data, uacomAssuntoCreated.CodAssunto))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, uacomAssuntoCreated)

}
