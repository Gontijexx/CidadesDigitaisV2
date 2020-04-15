package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
	STRUCT FATURA
=========================	*/

type Fatura struct {
	NumNF   uint32 `gorm:"primary_key;not null" json:"num_nf"`
	CodIbge uint32 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	DtNf    string `gorm:"default:null" json:"dt_nf"`
}

/*  =========================
	FUNCAO SALVAR FATURA
=========================  */

func (fatura *Fatura) SaveFatura(db *gorm.DB) (*Fatura, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&fatura).Error
	if err != nil {
		return &Fatura{}, err
	}

	return fatura, err
}

/*  =========================
	FUNCAO LISTAR FATURA POR ID
=========================  */

func (fatura *Fatura) FindFaturaByID(db *gorm.DB, numNF, codIbge uint32) (*Fatura, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Fatura{}).Where("num_nf = ? AND cod_ibge = ?", numNF, codIbge).Take(&fatura).Error
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
	FUNCAO EDITAR FATURA
========================   */

func (fatura *Fatura) UpdateFatura(db *gorm.DB, numNF, codIbge uint32) (*Fatura, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE fatura SET dt_nf = ? WHERE num_nf = ? AND cod_ibge = ?", fatura.DtNf, numNF, codIbge)
	if db.Error != nil {
		return &Fatura{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de suas chaves primarias
	err := db.Debug().Model(&Fatura{}).Where("num_nf = ? AND cod_ibge = ?", numNF, codIbge).Take(fatura).Error
	if err != nil {
		return &Fatura{}, err
	}

	return fatura, err
}

/*  =========================
	FUNCAO DELETAR FATURA
=========================  */

func (fatura *Fatura) DeleteFatura(db *gorm.DB, numNF, codIbge uint32) error {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Fatura{}).Where("num_nf = ? AND cod_ibge = ?", numNF, codIbge).Take(&Fatura{}).Delete(&Fatura{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("Fatura not found")
		}
		return db.Error
	}

	return db.Error
}
