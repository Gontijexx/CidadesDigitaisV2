package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR OTB NO BANCO DE DADOS
=========================  */

func (otb *OTB) SaveOTB(db *gorm.DB) (*OTB, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&otb).Error
	if err != nil {
		return &OTB{}, err
	}

	return otb, nil
}

/*  =========================
	FUNCAO LISTAR OTB POR ID
=========================  */

func (otb *OTB) FindOTBByID(db *gorm.DB, CodOTB uint64) (*OTB, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(OTB{}).Where("cod_otb = ?", CodOTB).Take(&otb).Error

	if err != nil {
		return &OTB{}, err
	}

	return otb, err
}

/*  =========================
	FUNCAO LISTAR TODAS OTB
=========================  */

func (otb *OTB) FindAllOTB(db *gorm.DB) (*[]OTB, error) {

	allOTB := []OTB{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&OTB{}).Find(&allOTB).Error
	if err != nil {
		return &[]OTB{}, err
	}
	return &allOTB, err
}

/*  =========================
	FUNCAO EDITAR OTB
=========================  */

func (otb *OTB) UpdateOTB(db *gorm.DB, CodOTB uint64) (*OTB, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&OTB{}).Where("cod_otb = ?", CodOTB).Updates(
		OTB{
			DtPgto: otb.DtPgto}).Error

	if db.Error != nil {
		return &OTB{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&OTB{}).Where("cod_otb = ?", CodOTB).Take(&otb).Error
	if err != nil {
		return &OTB{}, err
	}

	// retorna o elemento que foi alterado
	return otb, err
}

/*  =========================
	FUNCAO DELETAR OTB POR ID
=========================  */

func (otb *OTB) DeleteOTB(db *gorm.DB, CodOTB uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&OTB{}).Where("cod_otb = ?", CodOTB).Take(&OTB{}).Delete(&OTB{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("OTB not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
