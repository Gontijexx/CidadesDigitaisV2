package models

import (
	"CidadesDigitaisV2/api/config"
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

/*	=========================
		COMENTAR!!!!!
=========================	*/

type Usuario struct {
	Cod_usuario uint32 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_usuario"`
	Nome        string `gorm:"size:100;default:null" json:"nome" validate:"omitempty,alphanum"`
	Email       string `gorm:"size:45;default:null" json:"email" validate:"omitempty,email"`
	Status      bool   `gorm:"size:1;default:null" json:"status" `
	Login       string `gorm:"size:45;default:null" json:"login" validate:"alphanum"`
	Senha       string `gorm:"size:100;default:null" json:"senha" validate:"min=8"`
}

type Modulo struct {
	Cod_modulo  uint   `gorm:"primary_key;not null;size:11" json:"cod_modulo"`
	Categoria_1 string `gorm:"size:45;default:null" json:"cat_1"`
	Categoria_2 string `gorm:"size:45;default:null" json:"cat_2"`
	Categoria_3 string `gorm:"size:45;default:null" json:"cat_3"`
	Descricao   string `gorm:"size:200;default:null" json:"descricao"`
}

type Usuario_modulo struct {
	Cod_usuario uint32 `gorm:"foreingkey:Cod_usuario" `
	Cod_modulo  int64  `gorm:"foreingkey:Cod_modulo"`
}

func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerifyPassword(hashedSenha, senha string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}

func (u *Usuario) BeforeSave() error {

	hashedSenha, err := Hash(u.Senha)
	if err != nil {
		return err
	}
	u.Senha = string(hashedSenha)
	return nil
}

func (u *Usuario) Ready() {

	u.Cod_usuario = 0
	u.Status = true
	u.Nome = html.EscapeString(strings.TrimSpace(u.Nome))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Login = html.EscapeString(strings.TrimSpace(u.Login))
}

func (u *Usuario) SaveUser(db *gorm.DB) (*Usuario, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil

}

func (u *Usuario) FindUserByID(db *gorm.DB, uId uint32) (*Usuario, error) {

	err := db.Debug().Model(Usuario{}).Where("cod_usuario = ?", uId).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Usuario{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *Usuario) FindAllUsers(db *gorm.DB) (*[]Usuario, error) {

	usuario := []Usuario{}
	err := db.Debug().Model(&Usuario{}).Limit(100).Find(&usuario).Error
	if err != nil {
		return &[]Usuario{}, err
	}
	return &usuario, err
}

func (u *Usuario) UpdateAUser(db *gorm.DB, uId uint32) (*Usuario, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Printf("[FATAL] cannot HASH password, %v\n", err)
	}

	db = db.Debug().Model(&Usuario{}).Where("cod_usuario= ?", uId).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"senha":  u.Senha,
			"nome":   u.Nome,
			"email":  u.Email,
			"status": u.Status,
		},
	)
	if db.Error != nil {
		return &Usuario{}, db.Error
	}

	// This is the display the updated user
	err = db.Debug().Model(&Usuario{}).Where("cod_usuario = ?", uId).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}

	return u, nil

}

func (u *Usuario) DeleteAUser(db *gorm.DB, uId uint32) (int64, error) {

	db = db.Debug().Model(&Usuario{}).Where("cod_usuario = ?", uId).Take(&Usuario{}).Delete(&Usuario{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

func (u *Modulo) FindAllModulo(db *gorm.DB) (*[]Modulo, error) {

	modulo := []Modulo{}
	err := db.Debug().Model(&Modulo{}).Find(&modulo).Error
	if err != nil {
		return &[]Modulo{}, err
	}
	return &modulo, err
}

func (u *Usuario_modulo) CreateModulo(db *gorm.DB, uId uint32, mods interface{}) (*Usuario_modulo, error) {

	modulo := config.InterfaceSlice(mods)

	moduloInt := make([]float64, len(modulo))

	for i := range modulo {
		moduloInt[i] = modulo[i].(float64)
	}

	for _, v := range moduloInt {
		err := db.Debug().Raw("INSERT INTO usuario_modulo VALUE(?,?)", uId, v).Error
		if err != nil {
			return &Usuario_modulo{}, err
		}

	}

	return u, nil
}
