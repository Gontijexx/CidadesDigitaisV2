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
	FUNCAO ADICIONAR REAJUSTE
=========================  */

func (server *Server) AddReajuste(w http.ResponseWriter, r *http.Request) {
	
	//	Autorizacao de Modulo
	config.AuthMod(w, r, 14001)

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it coudn't read the 'body', %v\n", err))
	}

	//	Estrutura models.Reajuste{} "renomeada"
	reajuste := models.Reajuste{}

	//	Unmarshal analisa o JSON recebido e armazena na struct reajuste referenciada (&struct)
	err = json.Unmarshal(body, &reajuste)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(reajuste); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveReajuste eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	reajusteCreated, err := reajuste.SaveReajuste(server.DB)

	//	Retorna um erro caso nao seja possivel salvar reajuste no banco de dados
	//	Status 500
	if err != nil {

		formattedError := config.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it coudn't save in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, reajusteCreated.Ano_ref))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, reajusteCreated)

}

/*  =========================
	FUNCAO LISTAR REAJUSTE
=========================  */

func (server *Server) GetReajustes(w http.ResponseWriter, r *http.Request) {
	
	//	Autorizacao de Modulo
	config.AuthMod(w, r, 14002)

	reajuste := models.Reajuste{}

	//	reajustes armazena os dados buscados no banco de dados
	reajustes, err := reajuste.FindReajustes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] It couldn't find the result, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, reajustes)

}

/*  =========================
	FUNCAO LISTAR REAJUSTE POR ID
=========================  */

func (server *Server) GetReajuste(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 14002)

	//Vars retorna a rota das variaveis
	vars := mux.Vars(r)

	//interpreta  a string em uma base de (0, 2 to 36) e tamanho de (0 to 64)
	rId, err := strconv.ParseUint(vars["ano_ref"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	reajuste := models.Reajuste{}

	//vai utilizar o metodo para procurar o resultado de acordo com a chave
	reajusteGotten, err := reajuste.FindReajuste(server.DB, uint64(rId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//retorna um JSON indicando que funcionou corretamente
	responses.JSON(w, http.StatusOK, reajusteGotten)
}

/*  =========================
	FUNCAO ATUALIZAR REAJUSTE
=========================  */

func (server *Server) UpdateReajuste(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	config.AuthMod(w, r, 14003)

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	rIds armazena a chave primaria da tabela reajuste
	rId1, err := strconv.ParseUint(vars["ano_ref"], 10, 32)
	rId2, err := strconv.ParseUint(vars["cod_lote"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	reajuste := models.Reajuste{}

	err = json.Unmarshal(body, &reajuste)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(reajuste); err != nil {
		log.Printf("[WARN] invalid reajuste information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updatedReajuste recebe o novo reajuste, a que foi alterada
	updateReajuste, err := reajuste.UpdateReajuste(server.DB, uint64(rId1), uint64(rId2))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateReajuste)
}

/*  =========================
	FUNCAO DELETAR REAJUSTE
=========================  */

func (server *Server) DeleteReajuste(w http.ResponseWriter, r *http.Request) {

	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	reajuste := models.Reajuste{}

	//	rIDs armazena a chave primaria da tabela reajuste
	
	rID1, err := strconv.ParseUint(vars["ano_ref"], 10, 64)
	rID2, err := strconv.ParseUint(vars["cod_lote"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = reajuste.DeleteReajuste(server.DB, uint64(rID1), uint64(rID2))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", rID1, rID2))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}
