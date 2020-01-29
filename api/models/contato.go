package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		ALEXANDRE: IREI OLHAR AINDA
=========================	*/

/*  =========================
	FUNCAO SALVAR CONTATO
=========================  */

func (contato *Contato) SaveContato(db *gorm.DB) (*Contato, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&contato).Error
	if err != nil {
		return &Contato{}, err
	}
	return contato, nil

}

/*  =========================
	FUNCAO LISTAR CONTATO POR ID
=========================  */

func (contato *Contato) FindContatoByID(db *gorm.DB, contatoID uint64) (*Contato, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Contato{}).Where("cod_contato = ?", contatoID).Take(&contato).Error

	if err != nil {
		return &Contato{}, err
	}

	return contato, err
}

/*  =========================
	FUNCAO LISTAR CONTATOS
=========================  */

func (contato *Contato) FindAllContato(db *gorm.DB) (*[]Contato, error) {

	allContato := []Contato{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Contato{}).Limit(100).Find(&allContato).Error
	if err != nil {
		return &[]Contato{}, err
	}
	return &allContato, err
}

/*  =========================
	FUNCAO EDITAR CONTATO
=========================  */

func (contato *Contato) UpdateContato(db *gorm.DB, contatoID uint64) (*Contato, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Contato{}).Where("cod_contato = ?", contatoID).Updates(
		Contato{
			Nome:   contato.Nome,
			Email:  contato.Email,
			Funcao: contato.Funcao}).Error

	if db.Error != nil {
		return &Contato{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Contato{}).Where("cod_contato = ?", contatoID).Take(&contato).Error
	if err != nil {
		return &Contato{}, err
	}

	// retorna o elemento que foi alterado
	return contato, err
}

/*  =========================
	FUNCAO DELETAR ENTIDADE POR ID
=========================  */

func (contato *Contato) DeleteContato(db *gorm.DB, contatoID uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Contato{}).Where("cnpj = ?", contatoID).Take(&Contato{}).Delete(&Contato{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Contato not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
