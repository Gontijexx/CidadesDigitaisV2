package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR LOTE ITENS NO BANCO DE DADOS
=========================  */

func (loteItens *LoteItens) SaveLoteItens(db *gorm.DB) (*LoteItens, error) {

	//	Adicionao um novo elemento no banco de dados
	err := db.Debug().Create(&loteItens).Error
	if err != nil {
		return &LoteItens{}, err
	}

	return loteItens, nil
}

/*  =========================
	FUNCAO LISTAR LOTE ITENS POR ID
=========================  */

func (loteItens *LoteItens) FindLoteItensByID(db *gorm.DB, codLote, codItem, codTipoItem uint64) (*LoteItens, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", codLote, codItem, codTipoItem).Take(&loteItens).Error

	if err != nil {
		return &LoteItens{}, err
	}

	return loteItens, err
}

/*  =========================
	FUNCAO LISTAR TODOS LOTE ITENS
=========================  */

func (loteItens *LoteItens) FindAllLoteItens(db *gorm.DB) (*[]LoteItens, error) {

	allLoteItens := []LoteItens{}

	//	Busca todos elementos contidos no banco de dados

	err := db.Debug().Table("lote_itens").Select("itens.descricao, lote_itens.*").
		Joins("JOIN itens ON lote_itens.cod_item = itens.cod_item AND lote_itens.cod_tipo_item = itens.cod_tipo_item").Scan(&allLoteItens).Error

	if err != nil {
		return &[]LoteItens{}, err
	}
	return &allLoteItens, err
}

/*  =========================
	FUNCAO EDITAR LOTE ITENS
=========================  */

func (loteItens *LoteItens) UpdateLoteItens(db *gorm.DB, codLote, codItem, codTipoItem uint64) (*LoteItens, error) {

	err := db.Debug().Exec("UPDATE lote_itens SET preco = ? WHERE cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItens.Preco, codLote, codItem, codTipoItem).Error

	if db.Error != nil {
		return &LoteItens{}, db.Error
	}

	err = db.Debug().Model(&LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", codLote, codItem, codTipoItem).Take(&loteItens).Error
	if err != nil {
		return &LoteItens{}, err
	}

	return loteItens, err
}
