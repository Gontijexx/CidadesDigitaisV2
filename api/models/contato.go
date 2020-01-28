package models

import (
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

	entity := []Contato{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Contato{}).Limit(100).Find(&entity).Error
	if err != nil {
		return &[]Contato{}, err
	}
	return &entity, err
}

/*  =========================
	FUNCAO EDITAR CONTATO
=========================  */

func (contato *Contato) UpdateContato(db *gorm.DB, contatoID uint64) (*Contato, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Contato{}).Where("cod_contato = ?", contatoID).Take(&Contato{}).UpdateColumns(
		map[string]interface{}{
			"nome":   contato.Nome,
			"email":  contato.Email,
			"funcao": contato.Funcao,
		},
	)

	if db.Error != nil {
		return &Contato{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Contato{}).Where("cod_contato = ?", contatoID).Take(&contato).Error
	if err != nil {
		return &Contato{}, err
	}

	// retorna o elemento que foi alterado
	return contato, err
}
