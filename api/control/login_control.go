package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"CidadesDigitaisV2/api/Models"
	"CidadesDigitaisV2/api/auth"
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/responses"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := Models.Usuario{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Ready()

	token, err := server.SignIn(user.Login, user.Senha)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(login, password string) (string, error) {

	var err error

	user := Models.Usuario{}
	mod := Models.Usuario_modulo{}

	err = server.DB.Debug().Model(user).Where("login = ?", login).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = Models.VerifyPassword(user.Senha, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	err = server.DB.Debug().Select(mod.Cod_modulo).Where("Cod_usuario = ?", user.Cod_usuario).Find(&mod).Error
	if err != nil {
		return "", fmt.Errorf("[WARN] Cannot find users mod, because, %v\n", err)
	}

	return auth.CreateToken(user.Cod_usuario, mod.Cod_modulo)
}
