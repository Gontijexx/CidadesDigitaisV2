package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR TELEFONE
=========================  */

func (telefone *Telefone) SaveTelefone(db *gorm.DB) (*Telefone, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&telefone).Error
	if err != nil {
		return &Telefone{}, err
	}
	return telefone, nil

}

/*  =========================
	FUNCAO LISTAR TODAS TELEFONE
=========================  */

func (telefone *Telefone) FindAllTelefone(db *gorm.DB) (*[]Telefone, error) {

	allTelefone := []Telefone{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Telefone{}).Limit(100).Find(&allTelefone).Error
	if err != nil {
		return &[]Telefone{}, err
	}
	return &allTelefone, err
}

/*  =========================
	FUNCAO DELETAR TELEFONE POR ID
=========================  */

func (telefone *Telefone) DeleteTelefone(db *gorm.DB, codTelefone uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Telefone{}).Where("cod_telefone = ?", codTelefone).Take(&Telefone{}).Delete(&Telefone{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Telefone not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
