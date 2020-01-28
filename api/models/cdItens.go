package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		COMENTAR
=========================	*/

func (cdItens *CDItens) FindCDItensByID(db *gorm.DB, cdItensID1, cdItensID2, cdItensID3 uint64) (*CDItens, error) {

	err := db.Debug().Model(CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID1, cdItensID2, cdItensID3).Take(&cdItens).Error

	if err != nil {
		return &CDItens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CDItens{}, errors.New("Cd_itens Not Found")
	}

	return cdItens, err
}

func (cdItens *CDItens) UpdateCDItens(db *gorm.DB, cdItensID1, cdItensID2, cdItensID3 uint64) (*CDItens, error) {

	db = db.Debug().Model(&CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID1, cdItensID2, cdItensID3).Take(&cdItens).UpdateColumns(
		map[string]interface{}{
			"quantidade_previsto":          cdItens.Quantidade_previsto,
			"quantidade_projeto_executivo": cdItens.Quantidade_projeto_executivo,
			"quantidade_termo_instalacao":  cdItens.Quantidade_termo_instalacao,
		},
	)

	if db.Error != nil {
		return &CDItens{}, db.Error
	}

	err := db.Debug().Model(&CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID1, cdItensID2, cdItensID3).Take(&cdItens).Error
	if err != nil {
		return &CDItens{}, err
	}

	return cdItens, err
}

func (e *CDItens) FindAllCDItens(db *gorm.DB) (*[]CDItens, error) {
	cdItens := []CDItens{}
	err := db.Debug().Model(&CDItens{}).Find(&cdItens).Error
	if err != nil {
		return &[]CDItens{}, err
	}
	return &cdItens, err
}
