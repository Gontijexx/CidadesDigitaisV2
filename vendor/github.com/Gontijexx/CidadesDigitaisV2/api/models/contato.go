package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

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
	FUNCAO LISTAR TODOS CONTATOS
=========================  */

func (contato *Contato) FindAllContato(db *gorm.DB) (*[]Contato, error) {

	allContato := []Contato{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Contato{}).Find(&allContato).Error
	if err != nil {
		return &[]Contato{}, err
	}
	return &allContato, err
}

/*  =========================
	FUNCAO EDITAR CONTATO
=========================  */

func (contato *Contato) UpdateContato(db *gorm.DB, codContato uint64) (*Contato, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE contato SET cnpj =? , cod_ibge = ? , nome = ?, email = ?, funcao = ? WHERE cod_contato = ?", contato.Cnpj, contato.CodIbge, contato.Nome, contato.Email, contato.Funcao, codContato).Error

	if db.Error != nil {
		return &Contato{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Contato{}).Where("cod_contato = ?", codContato).Take(&contato).Error
	if err != nil {
		return &Contato{}, err
	}

	// retorna o elemento que foi alterado
	return contato, err
}

/*  =========================
	FUNCAO DELETAR ENTIDADE POR ID
=========================  */

func (contato *Contato) DeleteContato(db *gorm.DB, codContato uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Contato{}).Where("cod_contato = ?", codContato).Take(&Contato{}).Delete(&Contato{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Contato not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
