package control

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"CidadesDigitaisV2/api/auth"
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/models"
	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/validation"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
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

	user.Ready()

	token, err := server.SignIn(user.Login, user.Senha)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	if strings.Contains(token, "Error") {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(login, password string) (string, error) {

	var err error
	var CodMod []int64
	user := models.Usuario{}
	mods := models.Usuario_modulo{}

	err = server.DB.Debug().Model(user).Where("login = ?", login).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Senha, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	if user.Status == true {
		rows, err := server.DB.Debug().Raw("select cod_modulo from usuario_modulo where cod_usuario = ?", user.Cod_usuario).Rows()
		if err != nil {
			return "", err
		}
		for rows.Next() {

			err = rows.Scan(&mods.Cod_modulo)

			CodMod = append(CodMod, mods.Cod_modulo)

		}
		return auth.CreateToken(user.Cod_usuario, CodMod)
	} else {
		log.Printf("[FATAL] This user is disable,%v\n", user.Status)
		return "Error", err
	}

}
