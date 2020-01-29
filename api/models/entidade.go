package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ENTIDADE NO BANCO DE DADOS
=========================  */

func (entidade *Entidade) SaveEntidade(db *gorm.DB) (*Entidade, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}
	return entidade, nil

}

/*  =========================
	FUNCAO LISTAR ENTIDADE POR ID
=========================  */

func (entidade *Entidade) FindEntidadeByID(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error

	if err != nil {
		return &Entidade{}, err
	}

	return entidade, err
}

/*  =========================
	FUNCAO LISTAR TODAS ENTIDADE
=========================  */

func (entidade *Entidade) FindAllEntidade(db *gorm.DB) (*[]Entidade, error) {

	allEntidade := []Entidade{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Entidade{}).Limit(100).Find(&allEntidade).Error
	if err != nil {
		return &[]Entidade{}, err
	}
	return &allEntidade, err
}

/*  =========================
	FUNCAO EDITAR ENTIDADE
=========================  */

func (entidade *Entidade) UpdateEntidade(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Updates(
		Entidade{
			Nome:           entidade.Nome,
			Endereco:       entidade.Endereco,
			Numero:         entidade.Numero,
			Bairro:         entidade.Bairro,
			Cep:            entidade.Cep,
			Nome_municipio: entidade.Nome_municipio,
			Uf:             entidade.Uf,
			Observacao:     entidade.Observacao}).Error

	if db.Error != nil {
		return &Entidade{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}

	// retorna o elemento que foi alterado
	return entidade, err
}

/*  =========================
	FUNCAO DELETAR ENTIDADE POR ID
=========================  */

func (entidade *Entidade) DeleteEntidade(db *gorm.DB, entidadeID uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Take(&Entidade{}).Delete(&Entidade{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Entidade not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
