package Models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Cod_usuario uint32 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_usuario"`
	Nome        string `gorm:"size:100;default:null" json:"nome" validate: "alphanum":`
	Email       string `gorm:"size:45;default:null" json:"email" validate: "alphanum, email"`
	Status      bool   `gorm:"size:1;default:null" json:"status" validate: "alphanum"`
	Login       string `gorm:"size:45;default:null" json:"login" validate: "alphanum"`
	Senha       string `gorm:"size:45;default:null" json:"senha" validate: "alphanum, min=8"`
}

type Modulo struct {
	Cod_modulo  uint32   `gorm:"primary_key;not null;size:11" json:"cod_modulo"`
	Categoria_1 string   `gorm:"size:45;default:null" json:"cat_1"`
	Categoria_2 string   `gorm:"size:45;default:null" json:"cat_2"`
	Categoria_3 string   `gorm:"size:45;default:null" json:"cat_3"`
	Descricao   string   `gorm:"size:200;default:null" json:"descricao"`
}

type Usuario_modulo struct {
	Cod_usuario Usuario  `gorm:"foreingkey:Cod_usuario" `
	Cod_Modulo  Modulo `gorm:"foreingkey:Cod_modulo"`
}

func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerifyPassword(hashedSenha, senha string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Usuario) BeforeSave() error {

	hashedSenha, err := hash(u.Senha)
	if err != nil {
		return log.Printf("[WARN] invalid user information, because, %v\n", err)
	}
	u.Senha = string(hashedSenha)
}

func (u *Usuario) Ready() {

	u.Cod_usuario = 0
	u.Nome = html.EscapeString(strings.TrimSpace(u.Nome))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Login = html.EscapeString(strings.TrimSpace(u.Login))
}

func (u *Usuario) SaveUser(db *gorm.DB) (*Usuario, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &Usuario{}, log.Printf("[ERROR] cannot save user info, because, %v\n", err)
	}
	return u, nil

}

func (u *Usuario) FindUserByID(db *gorm.DB, uId uint32) (*Usuario, error) {

	err := db.Debug().Model(Usuario{}).Where("cod_usuario = ?", uId).Take(&u).Error
	if err != nil {
		return &Usuario{}, log.Printf("[ERROR] cannot find user by ID, because, %v\n", err)
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Usuario{}, errors.New("User Not Found"), log.Print("[WARN]User Not Found")
	}
	return u, err
}

func (u *Usuario) FindAllUsers(db *gorm.DB) (*[]Usuario, error) {

	usuarios := []Usuario{}
	err := db.Debug().Model(&Usuario{}).Limit(100).Find(&usuarios).Error
	if err != nil {
		return &Usuario{}, log.Printf("[ERROR] cannot find users because, %v\n", err)
	}
	return &usuarios, err
}

func (u *Usuario) UpdateAUser(db *gorm.DB, uId uint32) (*Usuario, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Printf("[FATAL] cannot HASH password, %v\n", err)
	}

	db = db.Debug().Model(&Usuario{}).Where("id = ?", uId).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"senha": u.Senha,
			"nome":  u.Nome,
			"email": u.Email,
		},
	)
	if db.Error != nil {
		return &Usuario{}, log.Printf("[ERROR] cannot update, %v\n", db.Error)
	}

	// This is the display the updated user
	err = db.Debug().Model(&Usuario{}).Where("cod_usuario = ?", uId).Take(&u).Error
	if err != nil {
		return &Usuario{}, log.Printf("[ERROR] cannot display updated user, %v\n", err)
	}

	return u, nil

}

func (u *Usuario) DeleteAUser(db *gorm.DB, uId uint32) (int64, error) {

	db = db.Debug().Model(&Usuario{}).Where("cod_usuario = ?", uId).Take(&Usuario{}).Delete(&Usuario{})

	if db.Error != nil {
		return 0, log.Printf("[ERROR] cannot delete, %v\n", db.Error)
	}

	return db.RowsAffected, nil
}
