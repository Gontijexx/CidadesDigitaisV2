package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR FATURA NO BANCO DE DADOS
=========================  */

func (fatura *Fatura) SaveFatura(db *gorm.DB) (*Fatura, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&fatura).Error
	if err != nil {
		return &Fatura{}, err
	}

	return fatura, nil
}

/*  =========================
	FUNCAO LISTAR FATURA POR ID
=========================  */

func (fatura *Fatura) FindFaturaByID(db *gorm.DB, numNF uint64) (*Fatura, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Fatura{}).Where("num_nf = ?", numNF).Take(&fatura).Error

	if err != nil {
		return &Fatura{}, err
	}

	return fatura, err
}

/*  =========================
	FUNCAO LISTAR TODAS FATURA
=========================  */

func (fatura *Fatura) FindAllFatura(db *gorm.DB) (*[]Fatura, error) {

	allFatura := []Fatura{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Fatura{}).Find(&allFatura).Error
	if err != nil {
		return &[]Fatura{}, err
	}
	return &allFatura, err
}

/*  =========================
	FUNCAO DELETAR FATURA POR ID
=========================  */

func (fatura *Fatura) DeleteFatura(db *gorm.DB, numNF uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Fatura{}).Where("num_nf = ?", numNF).Take(&Fatura{}).Delete(&Fatura{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Fatura not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
