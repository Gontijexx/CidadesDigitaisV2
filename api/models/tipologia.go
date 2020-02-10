package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR TIPOLOGIA NO BANCO DE DADOS
=========================  */

func (tipologia *Tipologia) SaveTipologia(db *gorm.DB) (*Tipologia, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&tipologia).Error
	if err != nil {
		return &Tipologia{}, err
	}

	return tipologia, nil
}

/*  =========================
	FUNCAO LISTAR TIPOLOGIA POR ID
=========================  */

func (tipologia *Tipologia) FindTipologiaByID(db *gorm.DB, codTipologia uint64) (*Tipologia, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Tipologia{}).Where("cod_tipologia = ?", codTipologia).Take(&tipologia).Error

	if err != nil {
		return &Tipologia{}, err
	}

	return tipologia, err
}

/*  =========================
	FUNCAO LISTAR TODAS TIPOLOGIA
=========================  */

func (tipologia *Tipologia) FindAllTipologia(db *gorm.DB) (*[]Tipologia, error) {

	allTipologia := []Tipologia{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Tipologia{}).Limit(100).Find(&allTipologia).Error
	if err != nil {
		return &[]Tipologia{}, err
	}

	return &allTipologia, err
}

/*  =========================
	FUNCAO EDITAR TIPOLOGIA
=========================  */

func (tipologia *Tipologia) UpdateTipologia(db *gorm.DB, codTipologia uint64) (*Tipologia, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Tipologia{}).Where("cod_tipologia = ?", codTipologia).Updates(
		Tipologia{
			Descricao: tipologia.Descricao}).Error

	if db.Error != nil {
		return &Tipologia{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Tipologia{}).Where("cod_tipologia = ?", codTipologia).Take(&tipologia).Error
	if err != nil {
		return &Tipologia{}, err
	}

	// retorna o elemento que foi alterado
	return tipologia, err
}

/*  =========================
	FUNCAO DELETAR TIPOLOGIA POR ID
=========================  */

func (tipologia *Tipologia) DeleteTipologia(db *gorm.DB, codTipologia uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Tipologia{}).Where("cod_tipologia = ?", codTipologia).Take(&Tipologia{}).Delete(&Tipologia{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Tipologia not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
