package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR TIPO_ITEM NO BANCO DE DADOS
=========================  */

func (tipoItem *TipoItem) SaveTipoItem(db *gorm.DB) (*TipoItem, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&tipoItem).Error
	if err != nil {
		return &TipoItem{}, err
	}

	return tipoItem, nil
}

/*  =========================
	FUNCAO LISTAR TIPO_ITEM POR ID
=========================  */

func (tipoItem *TipoItem) FindTipoItemByID(db *gorm.DB, codTipoItem uint64) (*TipoItem, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(TipoItem{}).Where("cod_tipo_item = ?", codTipoItem).Take(&tipoItem).Error

	if err != nil {
		return &TipoItem{}, err
	}

	return tipoItem, err
}

/*  =========================
	FUNCAO LISTAR TODAS TIPO_ITEM
=========================  */

func (tipoItem *TipoItem) FindAllTipoItem(db *gorm.DB) (*[]TipoItem, error) {

	allTipoItem := []TipoItem{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&TipoItem{}).Find(&allTipoItem).Error
	if err != nil {
		return &[]TipoItem{}, err
	}
	return &allTipoItem, err
}

/*  =========================
	FUNCAO EDITAR TIPO_ITEM
=========================  */

func (tipoItem *TipoItem) UpdateTipoItem(db *gorm.DB, codTipoItem uint64) (*TipoItem, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&TipoItem{}).Where("cod_tipo_item = ?", codTipoItem).Updates(
		TipoItem{
			Descricao: tipoItem.Descricao}).Error

	if db.Error != nil {
		return &TipoItem{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&TipoItem{}).Where("cod_tipo_item = ?", codTipoItem).Take(&tipoItem).Error
	if err != nil {
		return &TipoItem{}, err
	}

	// retorna o elemento que foi alterado
	return tipoItem, err
}

/*  =========================
	FUNCAO DELETAR TIPO_ITEM POR ID
=========================  */

func (tipoItem *TipoItem) DeleteTipoItem(db *gorm.DB, codTipoItem uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&TipoItem{}).Where("cod_tipo_item = ?", codTipoItem).Take(&TipoItem{}).Delete(&TipoItem{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("TipoItem not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
