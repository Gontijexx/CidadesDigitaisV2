package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func (cdItens *Cd_itens) FindCd_itensByID(db *gorm.DB, cdItensID, cdItensItem, cdItensTipo uint64) (*Cd_itens, error) {

	err := db.Debug().Model(Cd_itens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID, cdItensItem, cdItensTipo).Take(&cdItens).Error

	if err != nil {
		return &Cd_itens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cd_itens{}, errors.New("Cd_itens Not Found")
	}

	return cdItens, err
}

func (cdItens *Cd_itens) UpdateCd_itens(db *gorm.DB, cdItensID, cdItensItem, cdItensTipo uint64) (*Cd_itens, error) {

	db = db.Debug().Model(&Cd_itens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID, cdItensItem, cdItensTipo).Take(&cdItens).UpdateColumns(
		map[string]interface{}{
			"quantidade_previsto":          cdItens.Quantidade_previsto,
			"quantidade_projeto_executivo": cdItens.Quantidade_projeto_executivo,
			"quantidade_termo_instalacao":  cdItens.Quantidade_termo_instalacao,
		},
	)

	if db.Error != nil {
		return &Cd_itens{}, db.Error
	}

	err := db.Debug().Model(&Cd_itens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID, cdItensItem, cdItensTipo).Take(&cdItens).Error
	if err != nil {
		return &Cd_itens{}, err
	}

	return cdItens, err
}

func (e *Cd_itens) FindAllCd_itens(db *gorm.DB) (*[]Cd_itens, error) {
	cdItens := []Cd_itens{}
	err := db.Debug().Model(&Cd_itens{}).Find(&cdItens).Error
	if err != nil {
		return &[]Cd_itens{}, err
	}
	return &cdItens, err
}
