package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ASSUNTO NO BANCO DE DADOS
=========================  */

func (assunto *Assunto) SaveAssunto(db *gorm.DB) (*Assunto, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&assunto).Error
	if err != nil {
		return &Assunto{}, err
	}
	return assunto, nil

}

/*  =========================
	FUNCAO LISTAR ASSUNTO POR ID
=========================  */

func (assunto *Assunto) FindAssuntoByID(db *gorm.DB, codAssunto uint64) (*Assunto, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Assunto{}).Where("cod_assunto = ?", codAssunto).Take(&assunto).Error

	if err != nil {
		return &Assunto{}, err
	}

	return assunto, err
}

/*  =========================
	FUNCAO LISTAR TODOS ASSUNTOS
=========================  */

func (assunto *Assunto) FindAllAssunto(db *gorm.DB) (*[]Assunto, error) {

	allAssunto := []Assunto{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Assunto{}).Limit(100).Find(&allAssunto).Error
	if err != nil {
		return &[]Assunto{}, err
	}
	return &allAssunto, err
}

/*  =========================
	FUNCAO EDITAR ASSUNTO
=========================  */

func (assunto *Assunto) UpdateAssunto(db *gorm.DB, codAssunto uint64) (*Assunto, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Assunto{}).Where("cod_assunto = ?", codAssunto).Updates(
		Assunto{
			Descricao: assunto.Descricao}).Error

	if db.Error != nil {
		return &Assunto{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Assunto{}).Where("cod_assunto = ?", codAssunto).Take(&assunto).Error
	if err != nil {
		return &Assunto{}, err
	}

	// retorna o elemento que foi alterado
	return assunto, err
}

/*  =========================
	FUNCAO DELETAR ASSUNTO POR ID
=========================  */

func (assunto *Assunto) DeleteAssunto(db *gorm.DB, codAssunto uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Assunto{}).Where("cod_assunto = ?", codAssunto).Take(&Assunto{}).Delete(&Assunto{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Assunto not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
