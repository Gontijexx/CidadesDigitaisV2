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

/*	=========================
		COMENTAR!!! E ATUALIZAR MENSAGENS DE ERRO
=========================	*/

/*  =========================
	FUNCAO ADICIONAR USUARIO
=========================  */
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11001)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.Usuario{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(user); err != nil {
		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	NewUser := user.Login
	err = server.DB.Debug().Model(user).Where("login = ?", NewUser).Take(&user).Error
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"Error": "Existent Login"}`))
		return
	}

	user.Ready()
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := config.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.Cod_usuario))
	responses.JSON(w, http.StatusCreated, userCreated)

}

/*  =========================
	FUNCAO LISTAR TODOS
		USUARIOS
=========================  */
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}

	user := models.Usuario{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

/*  =========================
	FUNCAO LISTAR UM USUARIO
=========================  */
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11002)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	vars := mux.Vars(r)
	uId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.Usuario{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

/*  =========================
	FUNCAO PARA ATUALIZAR
		 USUARIO
=========================  */
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 11003)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.Usuario{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err = validation.Validator.Struct(user); err != nil {
		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user.Ready()

	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) AddModulo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	IDUser, err := strconv.ParseUint(vars["id2"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	uID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userModulo := models.Usuario_modulo{}

	err = json.Unmarshal(body, &userModulo)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(userModulo); err != nil {
		log.Printf("[WARN] invalid user_modulo information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	userMod, err := userModulo.CreateModulo(server.DB, uint32(uID), userModulo.Cod_modulo)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusCreated, userMod)
	return
}

/*  ============================
    FUNCAO PARA DELETAR USUARIOS
=============================
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {


	vars := mux.Vars(r)

	user := models.Usuario{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
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
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
*/
