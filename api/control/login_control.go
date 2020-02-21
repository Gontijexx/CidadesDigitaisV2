package control

import (
	"encoding/json"
	"fmt"

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

/*	=========================
		COMENTAR
=========================	*/

/*  =========================
	FUNCAO DE LOGIN
=========================  */

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	usuario := models.Usuario{}

	err = json.Unmarshal(body, &usuario)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = validation.Validator.Struct(usuario); err != nil {
		log.Printf("[WARN] invalid usuario information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	usuario.Prepare()

	token, err := server.SignIn(usuario.Login, usuario.Senha)

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

	var CodigoModulo []int64
	usuario := models.Usuario{}
	modulo := models.UsuarioModulo{}

	err := server.DB.Debug().Model(usuario).Where("login = ?", login).Take(&usuario).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(usuario.Senha, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	if usuario.Status == true {
		rows, err := server.DB.Debug().Raw("SELECT cod_modulo FROM usuario_modulo WHERE cod_usuario = ?", usuario.CodUsuario).Rows()
		if err != nil {
			return "", err
		}
		for rows.Next() {

			err = rows.Scan(&modulo.CodModulo)

			CodigoModulo = append(CodigoModulo, modulo.CodModulo)

		}
		fmt.Printf("eu so codmod: %v", CodigoModulo)
		return auth.CreateToken(usuario.CodUsuario, CodigoModulo)
	} else {
		log.Printf("[FATAL] This usuario is disable,%v\n", usuario.Status)
		return "Error", err
	}

}
