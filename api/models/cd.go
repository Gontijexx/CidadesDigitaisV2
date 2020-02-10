package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR CD
=========================  */

func (cd *CD) SaveCD(db *gorm.DB) (*CD, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&cd).Error
	if err != nil {
		return &CD{}, err
	}
	return cd, nil

}

/*  =========================
	FUNCAO LISTAR CD POR ID
=========================  */

func (cd *CD) FindCDByID(db *gorm.DB, codIbge uint64) (*CD, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(CD{}).Where("cod_ibge = ?", codIbge).Take(&cd).Error

	if err != nil {
		return &CD{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CD{}, errors.New("Cd Not Found")
	}

	return cd, err
}

/*  =========================
	FUNCAO LISTAR CD
=========================  */

func (cd *CD) FindAllCD(db *gorm.DB) (*[]CD, error) {

	allCD := []CD{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&CD{}).Limit(100).Find(&allCD).Error
	if err != nil {
		return &[]CD{}, err
	}
	return &allCD, err
}

/*  =========================
	FUNCAO EDITAR CD
=========================  */

func (cd *CD) UpdateCD(db *gorm.DB, codIbge uint64) (*CD, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&CD{}).Where("cod_ibge = ?", codIbge).Updates(
		CD{
			CodLote: cd.CodLote,
			OsPe:    cd.OsPe,
			DataPe:  cd.DataPe,
			OsImp:   cd.OsImp,
			DataImp: cd.DataImp}).Error

	if err != nil {
		return &CD{}, err
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&CD{}).Where("cod_ibge = ?", codIbge).Take(&cd).Error
	if err != nil {
		return &CD{}, err
	}

	// retorna o elemento que foi alterado
	return cd, err
}

/*  =========================
	FUNCAO DELETAR CD POR ID
=========================  */

func (cd *CD) DeleteCD(db *gorm.DB, codIbge uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&CD{}).Where("cod_ibge = ?", codIbge).Take(&CD{}).Delete(&CD{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CD not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
