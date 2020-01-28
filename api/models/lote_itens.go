package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		COMENTAR
=========================	*/

func (loteItens *LoteItens) FindLoteItensByID(db *gorm.DB, loteItensID1, loteItensID2, loteItensID3 uint64) (*LoteItens, error) {

	err := db.Debug().Model(LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItensID1, loteItensID2, loteItensID3).Take(&loteItens).Error

	if err != nil {
		return &LoteItens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &LoteItens{}, errors.New("Lote_itens Not Found")
	}

	return loteItens, err
}

func (loteItens *LoteItens) UpdateLoteItens(db *gorm.DB, loteItensID1, loteItensID2, loteItensID3 uint64) (*LoteItens, error) {

	db = db.Debug().Model(&LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItensID1, loteItensID2, loteItensID3).Take(&loteItens).UpdateColumns(
		map[string]interface{}{
			"preco": loteItens.Preco,
		},
	)

	if db.Error != nil {
		return &LoteItens{}, db.Error
	}

	err := db.Debug().Model(&LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteItensID1, loteItensID2, loteItensID3).Take(&loteItens).Error
	if err != nil {
		return &LoteItens{}, err
	}

	return loteItens, err
}

func (loteItens *LoteItens) FindAllLoteItens(db *gorm.DB) (*[]LoteItens, error) {
	allLoteItens := []LoteItens{}
	err := db.Debug().Model(&LoteItens{}).Find(&allLoteItens).Error
	if err != nil {
		return &[]LoteItens{}, err
	}
	return &allLoteItens, err
}
