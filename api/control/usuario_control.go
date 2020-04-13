package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"CidadesDigitaisV2/api/auth"
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/models"
	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/validation"

	"github.com/gorilla/mux"
)

/*  =========================
	FUNCAO ADICIONAR USUARIO
=========================  */
func (server *Server) CreateUsuario(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11001)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	O metodo ReadAll le toda a request ate encontrar algum erro, se nao encontrar erro o leitura para em EOF
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the body, %v\n", err))
	}

	//	Estrutura models.Usuario{} "renomeada"
	usuario := models.Usuario{}

	//	Unmarshal analisa o JSON recebido e armazena na struct usuario referenciada (&struct)
	err = json.Unmarshal(body, &usuario)

	//	Se ocorrer algum tipo de erro retorna-se o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(usuario); err != nil {
		log.Printf("[WARN] invalid usuario information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	newUsuario recebe o login a ser cadastrado
	newUsuario := usuario.Login

	//	Verifica se o login ja esta em uso
	err = usuario.VerifyLogin(server.DB, newUsuario)

	//	Tratamento do err, caso 'err != nil' pode-se criar o usuario
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"Error": "Existent Login"}`))
		return
	}

	//	Prepara algumas informacoes a serem salvas no banco de dados
	usuario.Prepare()

	//	SaveUsuario eh o metodo que faz a conexao com banco de dados e salva os dados recebidos
	usuarioCreated, err := usuario.SaveUsuario(server.DB)

	//	Retorna um erro caso nao seja possivel salvar entidade no banco de dados
	//	Status 500
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, usuarioCreated.CodUsuario))

	//	Ao final retorna o Status 201 e o JSON da struct que foi criada
	responses.JSON(w, http.StatusCreated, usuarioCreated)

}

/*  =========================
	FUNCAO LISTAR UM USUARIO POR ID
=========================  */

func (server *Server) GetUsuarioByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codUsuario armazena a chave primaria da tabela usuario
	codUsuario, err := strconv.ParseUint(vars["cod_usuario"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	usuario := models.Usuario{}

	//	usuarioGotten recebe o dado buscado no banco de dados
	usuarioGotten, err := usuario.FindUsuarioByID(server.DB, uint32(codUsuario))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, usuarioGotten)

}

/*  =========================
	FUNCAO LISTAR TODOS USUARIOS
=========================  */

func (server *Server) GetAllUsuario(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	usuario := models.Usuario{}

	//	allUsuario armazena os dados buscados no banco de dados
	allUsuario, err := usuario.FindAllUsuario(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allUsuario)
}

/*  =========================
	FUNCAO PARA ATUALIZAR USUARIO
=========================  */

func (server *Server) UpdateUsuario(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	codUsuario armazena a chave primaria da tabela usuario
	codUsuario, err := strconv.ParseUint(vars["cod_usuario"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	usuario := models.Usuario{}

	err = json.Unmarshal(body, &usuario)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(usuario); err != nil {
		log.Printf("[WARN] invalid usuario information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	Extrai o cod_usuario do body
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//	Verifica se o cod_usuario extraido do token eh igual ao extraido do handler
	if tokenID != uint32(codUsuario) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	//	updateUsuario recebe o novo usuario, o que foi alterado
	updatedUsuario, err := usuario.UpdateUsuario(server.DB, uint32(usuario.CodUsuario))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updatedUsuario)

}

/*  ============================
    FUNCAO PARA DELETAR USUARIOS
=============================
func (server *Server) DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	usuario := models.Usuario{}

	uid, err := strconv.ParseUint(vars["cod_usuario"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = usuario.DeleteUsuario(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
*/

/*  =========================
	FUNCAO LISTAR TODOS MODULOS
=========================  */

func (server *Server) GetAllModulo(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	modulo := models.Modulo{}

	//	allModulo armazena os dados buscados no banco de dados
	allModulo, err := modulo.FindAllModulo(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in datebase, %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allModulo)
}

/*  =========================
	FUNCAO PARA ADICIONAR MODULO
=========================  */

func (server *Server) AddModulo(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11003)
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

	//	Declara usarioModulo como um Array da struct UsuarioModulo
	var usuarioModulo []models.UsuarioModulo

	//	Unmarshal analisa o JSON recebido e armazena na struct usuarioModulo referenciada
	err = json.Unmarshal(body, &usuarioModulo)

	//	Caso err seja diferente de nil retorna o Status 422 mais o erro ocorrido
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	//	range le o tamanho do Array usuarioModulo
	//	data recebe os dados de cada posicao do array
	for _, data := range usuarioModulo {

		//	SaveModulo salva as informacoes contidas em 'data' no banco de dados
		//	usuarioMod recebe a struct salva no banco de dados
		usuarioMod, err := data.SaveModulo(server.DB)

		//	Retorna o erro caso nao seja possivel salvar o modulo no banco de dados
		if err != nil {
			formattedError := config.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't save in database, %v\n", formattedError))
			return
		}

		//	A cada posicao do Array usuarioModulo salva no banco de dados retorna o JSON da struct
		fmt.Printf("CodModulo adicionado: %v\n", usuarioMod.CodModulo)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, usuarioModulo))

	//	Retorna o Status 201 e o JSON do Array adicionado
	responses.JSON(w, http.StatusCreated, usuarioModulo)

}
