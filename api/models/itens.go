package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ITENS NO BANCO DE DADOS
=========================  */

func (itens *Itens) SaveItens(db *gorm.DB) (*Itens, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&itens).Error
	if err != nil {
		return &Itens{}, err
	}

	return itens, nil
}

/*  =========================
	FUNCAO LISTAR ITENS POR ID
=========================  */

func (itens *Itens) FindItensByID(db *gorm.DB, codItem, codTipoItem uint64) (*Itens, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Itens{}).Where("cod_item = ? AND cod_tipo_item = ?", codItem, codTipoItem).Take(&itens).Error

	if err != nil {
		return &Itens{}, err
	}

	return itens, err
}

/*  =========================
	FUNCAO LISTAR TODAS ITENS
=========================  */

func (itens *Itens) FindAllItens(db *gorm.DB) (*[]Itens, error) {

	allItens := []Itens{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Itens{}).Find(&allItens).Error
	if err != nil {
		return &[]Itens{}, err
	}
	return &allItens, err
}

/*  =========================
	FUNCAO EDITAR ITENS
=========================  */

func (itens *Itens) UpdateItens(db *gorm.DB, codItem, codTipoItem uint64) (*Itens, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Itens{}).Where("cod_item = ? AND cod_tipo_item", codItem, codTipoItem).Updates(
		Itens{
			Descricao: itens.Descricao,
			Unidade:   itens.Unidade}).Error

	if db.Error != nil {
		return &Itens{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Itens{}).Where("cod_item = ? AND cod_tipo_item", codItem, codTipoItem).Take(&itens).Error
	if err != nil {
		return &Itens{}, err
	}

	// retorna o elemento que foi alterado
	return itens, err
}

/*  =========================
	FUNCAO DELETAR ITENS POR ID
=========================  */

func (itens *Itens) DeleteItens(db *gorm.DB, codItem, codTipoItem uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Itens{}).Where("cod_item = ? AND cod_tipo_item", codItem, codTipoItem).Take(&Itens{}).Delete(&Itens{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Itens not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
