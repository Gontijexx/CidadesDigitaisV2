package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PROCESSO NO BANCO DE DADOS
=========================  */

func (processo *Processo) SaveProcesso(db *gorm.DB) (*Processo, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&processo).Error
	if err != nil {
		return &Processo{}, err
	}
	return processo, nil

}

/*  =========================
	FUNCAO LISTAR PROCESSO POR ID
=========================  */

func (processo *Processo) FindProcessoByID(db *gorm.DB, codProcesso uint64) (*Processo, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Processo{}).Where("cod_processo = ?", codProcesso).Take(&processo).Error

	if err != nil {
		return &Processo{}, err
	}

	return processo, err
}

/*  =========================
	FUNCAO LISTAR TODAS PROCESSO
=========================  */

func (processo *Processo) FindAllProcesso(db *gorm.DB) (*[]Processo, error) {

	allProcesso := []Processo{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Processo{}).Limit(100).Find(&allProcesso).Error
	if err != nil {
		return &[]Processo{}, err
	}
	return &allProcesso, err
}

/*  =========================
	FUNCAO EDITAR PROCESSO
=========================  */

func (processo *Processo) UpdateProcesso(db *gorm.DB, codProcesso uint64) (*Processo, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Processo{}).Where("cod_processo = ?", codProcesso).Updates(
		Processo{
			Descricao: processo.Descricao}).Error

	if db.Error != nil {
		return &Processo{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Processo{}).Where("cod_processo = ?", codProcesso).Take(&processo).Error
	if err != nil {
		return &Processo{}, err
	}

	// retorna o elemento que foi alterado
	return processo, err
}

/*  =========================
	FUNCAO DELETAR PROCESSO POR ID
=========================  */

func (processo *Processo) DeleteProcesso(db *gorm.DB, codProcesso uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Processo{}).Where("cod_processo = ?", codProcesso).Take(&Processo{}).Delete(&Processo{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Processo not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
