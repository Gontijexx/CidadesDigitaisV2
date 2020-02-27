package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

/*  =========================
	TABELA USUARIO
=========================  */

type Usuario struct {
	CodUsuario uint32 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_usuario"`
	Nome       string `gorm:"size:100;default:null" json:"nome"`
	Email      string `gorm:"size:45;default:null" json:"email" validate:"omitempty,email"`
	Status     bool   `gorm:"size:1;default:null" json:"status" `
	Login      string `gorm:"size:45;default:null" json:"login" validate:"alphanum"`
	Senha      string `gorm:"size:100;default:null" json:"senha" validate:"min=8"`
}

/*  =========================
	TABELA MODULO
=========================  */

type Modulo struct {
	CodModulo  uint   `gorm:"primary_key;not null;size:11" json:"cod_modulo"`
	Categoria1 string `gorm:"size:45;default:null" json:"categoria_1"`
	Categoria2 string `gorm:"size:45;default:null" json:"categoria_2"`
	Categoria3 string `gorm:"size:45;default:null" json:"categoria_3"`
	Descricao  string `gorm:"size:200;default:null" json:"descricao"`
}

/*  =========================
	TABELA USUARIO_MODULO
=========================  */

type UsuarioModulo struct {
	CodUsuario uint32 `gorm:"foreingkey:CodUsuario" json:"cod_usuario" validate:"required"`
	CodModulo  int64  `gorm:"foreingkey:CodModulo" json:"cod_modulo" validate:"required"`
}

/*	=========================
		COMENTAR!!!!!
=========================	*/

func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

/*	=========================
		COMENTAR!!!!!
=========================	*/

func VerifyPassword(hashedSenha, senha string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}

/*	=========================
		COMENTAR!!!!!
=========================	*/

func (usuario *Usuario) BeforeSave() error {

	hashedSenha, err := Hash(usuario.Senha)
	if err != nil {
		return err
	}
	usuario.Senha = string(hashedSenha)
	return nil
}

/*	=========================
		COMENTAR!!!!!
=========================	*/

func (usuario *Usuario) Prepare() {

	usuario.CodUsuario = 0
	usuario.Nome = html.EscapeString(strings.TrimSpace(usuario.Nome))
	usuario.Email = html.EscapeString(strings.TrimSpace(usuario.Email))
	usuario.Status = true
	usuario.Login = html.EscapeString(strings.TrimSpace(usuario.Login))
}

/*  =========================
	FUNCAO SALVAR USUARIO NO BANCO DE DADOS
=========================  */

func (usuario *Usuario) SaveUsuario(db *gorm.DB) (*Usuario, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&usuario).Error
	if err != nil {
		return &Usuario{}, err
	}

	return usuario, nil

}

/*  =========================
	FUNCAO LISTAR USUARIO POR ID
=========================  */

func (usuario *Usuario) FindUsuarioByID(db *gorm.DB, codUsuario uint32) (*Usuario, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Usuario{}).Where("cod_usuario = ?", codUsuario).Take(&usuario).Error
	if err != nil {
		return &Usuario{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Usuario{}, errors.New("Usuario Not Found")
	}
	return usuario, err
}

/*  =========================
	FUNCAO LISTAR TODOS USUARIO
=========================  */

func (usuario *Usuario) FindAllUsuario(db *gorm.DB) (*[]Usuario, error) {

	allUsuario := []Usuario{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Usuario{}).Find(&allUsuario).Error
	if err != nil {
		return &[]Usuario{}, err
	}
	return &allUsuario, err
}

/*  =========================
	FUNCAO EDITAR USUARIO
=========================  */

func (usuario *Usuario) UpdateUsuario(db *gorm.DB, codUsuario uint32) (*Usuario, error) {

	// To hash the password
	err := usuario.BeforeSave()
	if err != nil {
		log.Printf("[FATAL] cannot HASH password, %v\n", err)
	}

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Usuario{}).Where("cod_usuario= ?", codUsuario).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"nome":   usuario.Nome,
			"email":  usuario.Email,
			"status": usuario.Status,
			"senha":  usuario.Senha,
		},
	)
	if db.Error != nil {
		return &Usuario{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Usuario{}).Where("cod_usuario = ?", VerifyPassword).Take(&usuario).Error
	if err != nil {
		return &Usuario{}, err
	}

	// retorna o elemento que foi alterado
	return usuario, nil

}

/*  =========================
	FUNCAO DELETAR USUARIO POR ID
=========================  */

func (usuario *Usuario) DeleteUsuario(db *gorm.DB, codUsuario uint32) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Usuario{}).Where("cod_usuario = ?", codUsuario).Take(&Usuario{}).Delete(&Usuario{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

/*	=========================
		COMENTAR!!!!!
=========================	*/

func (usuario *Modulo) FindAllModulo(db *gorm.DB) (*[]Modulo, error) {

	modulo := []Modulo{}
	err := db.Debug().Model(&Modulo{}).Find(&modulo).Error
	if err != nil {
		return &[]Modulo{}, err
	}
	return &modulo, err
}

/*	=========================
		COMENTAR!!!!!
=========================	*/

func (usuarioModulo *UsuarioModulo) SaveModulo(db *gorm.DB) (*UsuarioModulo, error) {

	err := db.Debug().Create(&usuarioModulo).Error
	if err != nil {
		return &UsuarioModulo{}, err
	}
	return usuarioModulo, nil
}
