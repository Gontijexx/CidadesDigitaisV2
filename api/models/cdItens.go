package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	STRUCT CD ITENS
=========================  */

type CDItens struct {
	CodIbge                    uint32  `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	CodItem                    uint32  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem                uint32  `gorm:"primary_key;foreign_key:CodTipoItem;not null" json:"cod_tipo_item"`
	QuantidadePrevisto         uint32  `gorm:"default:null" json:"quantidade_previsto"`
	QuantidadeProjetoExecutivo float32 `gorm:"default:null" json:"quantidade_projeto_executivo"`
	QuantidadeTermoInstalacao  float32 `gorm:"default:null" json:"quantidade_termo_instalacao"`
}

/*  =========================
	FUNCAO LISTAR CD ITENS POR ID
=========================  */

func (cdItens *CDItens) FindCDItensByID(db *gorm.DB, codIbge, codItem, codTipoItem uint32) (*CDItens, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", codIbge, codItem, codTipoItem).Take(&cdItens).Error
	if err != nil {
		return &CDItens{}, err
	}

	return cdItens, err
}

/*  =========================
	FUNCAO LISTAR TODOS CD ITENS
=========================  */

func (cdItens *CDItens) FindAllCDItens(db *gorm.DB) (*[]CDItens, error) {

	allCDItens := []CDItens{}

	//	Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&CDItens{}).Find(&allCDItens).Error
	if err != nil {
		return &[]CDItens{}, err
	}

	return &allCDItens, err
}

/*  =========================
	FUNCAO EDITAR CD ITENS
=========================  */

func (cdItens *CDItens) UpdateCDItens(db *gorm.DB, codIbge, codItem, codTipoItem uint32) (*CDItens, error) {

	err := db.Debug().Exec("UPDATE cd_itens SET quantidade_previsto = ?, quantidade_projeto_executivo = ?, quantidade_termo_instalacao = ? WHERE cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItens.QuantidadePrevisto, cdItens.QuantidadeProjetoExecutivo, cdItens.QuantidadeTermoInstalacao, codIbge, codItem, codTipoItem).Error
	if err != nil {
		return &CDItens{}, err
	}

	err = db.Debug().Model(&CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", codIbge, codItem, codTipoItem).Take(&cdItens).Error
	if err != nil {
		return &CDItens{}, err
	}

	return cdItens, err
}
