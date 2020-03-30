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

/*  =========================
	FUNCAO ADICIONAR MUNICIPIO
=========================  */

func (server *Server) CreateMunicipio(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 24001)
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

	//	Estrutura models.Municipio{} "renomeada"
	municipio := models.Municipio{}

	//	Unmarshal analisa o JSON recebido e armazena na struct municipio referenciada (&struct)
	err = json.Unmarshal(body, &municipio)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(municipio); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	Validacao de CNPJ
	result := validation.ValidationCNPJ(municipio.Cnpj)
	if result == false {
		log.Printf("[FATAL] invalid CNPJ!")
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	SaveMunicipio eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	municipioCreated, err := municipio.SaveMunicipio(server.DB)

	/*	Retorna um erro caso nao seja possivel salvar municipio no banco de dados
		Status 500	*/
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, municipioCreated.CodIbge))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, municipioCreated)

}

/*  =========================
	FUNCAO LISTAR MUNICIPIO POR ID
=========================  */

func (server *Server) GetMunicipioByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 24002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codIbge armazena a chave primaria da tabela municipio
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	municipio := models.Municipio{}

	//	municipioGotten recebe o dado buscado no banco de dados
	municipioGotten, err := municipio.FindMunicipioByID(server.DB, codIbge)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, municipioGotten)

}

/*  =========================
	FUNCAO LISTAR TODOS MUNICIPIO
=========================  */

func (server *Server) GetAllMunicipio(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 24002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	municipio := models.Municipio{}

	//	allMunicipio armazena os dados buscados no banco de dados
	allMunicipio, err := municipio.FindAllMunicipio(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allMunicipio)
}

/*  =========================
	FUNCAO EDITAR MUNICIPIO
=========================  */

func (server *Server) UpdateMunicipio(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 24003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codIbge armazena a chave primaria da tabela municipio
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

	municipio := models.Municipio{}

	err = json.Unmarshal(body, &municipio)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(municipio); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateMunicipio recebe o novo municipio, a que foi alterada
	updateMunicipio, err := municipio.UpdateMunicipio(server.DB, codIbge)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateMunicipio)
}

/*  =========================
	FUNCAO DELETAR MUNICIPIO
=========================  */

func (server *Server) DeleteMunicipio(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo, apenas quem tem permicao de edit pode deletar
	err := config.AuthMod(w, r, 24003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	// Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	municipio := models.Municipio{}

	//	codIbge armazena a chave primaria da tabela municipio
	codIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	/* 	Para o caso da funcao 'delete' apenas o erro nos eh necessario
	Caso nao seja possivel deletar o dado especificado tratamos o erro*/
	_, err = municipio.DeleteMunicipio(server.DB, codIbge)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't delete in database , %v\n", formattedError))
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", codIbge))

	//	Retorna o Status 204, indicando que a informacao foi deletada
	responses.JSON(w, http.StatusNoContent, "")
}

/*  =========================
	FUNCAO LISTAR MUNICIPIO.CODIBGE E MUNICIPIO.NOMEMUNICIPIO
=========================  */

func (server *Server) GetMunicipioIDandNomeMunicipio(w http.ResponseWriter, r *http.Request) {

	municipio := models.Municipio{}

	//	municipioGotten recebe o dado buscado no banco de dados
	municipioGotten, err := municipio.FindMunicipioIDandNomeMunicipio(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	bytes, _ := json.Marshal(municipioGotten)

	w.Write(bytes)
}
