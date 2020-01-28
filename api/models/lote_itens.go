package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		COMENTAR
=========================	*/

func (loteItens *Lote_itens) FindLote_itensByID(db *gorm.DB, loteItensID1, loteItensID2, loteItensID3 uint64) (*Lote_itens, error) {

	err := db.Debug().Model(Lote_itens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItensID1, loteItensID2, loteItensID3).Take(&loteItens).Error

	if err != nil {
		return &Lote_itens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Lote_itens{}, errors.New("Lote_itens Not Found")
	}

	return loteItens, err
}

func (loteItens *Lote_itens) UpdateLote_itens(db *gorm.DB, loteItensID1, loteItensID2, loteItensID3 uint64) (*Lote_itens, error) {

	db = db.Debug().Model(&Lote_itens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItensID1, loteItensID2, loteItensID3).Take(&loteItens).UpdateColumns(
		map[string]interface{}{
			"preco": loteItens.Preco,
		},
	)

	if db.Error != nil {
		return &Lote_itens{}, db.Error
	}

	err := db.Debug().Model(&Lote_itens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItensID1, loteItensID2, loteItensID3).Take(&loteItens).Error
	if err != nil {
		return &Lote_itens{}, err
	}

	return loteItens, err
}

func (e *Lote_itens) FindAllLote_itens(db *gorm.DB) (*[]Lote_itens, error) {
	loteItens := []Lote_itens{}
	err := db.Debug().Model(&Lote_itens{}).Find(&loteItens).Error
	if err != nil {
		return &[]Lote_itens{}, err
	}
	return &loteItens, err
}
