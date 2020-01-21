package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func (lote *Lote) SaveLote(db *gorm.DB) (*Lote, error) {

	err := db.Debug().Create(&lote).Error
	if err != nil {
		return &Lote{}, err
	}
	return lote, nil

}

func (lote *Lote) FindLoteByID(db *gorm.DB, loteID uint64) (*Lote, error) {

	err := db.Debug().Model(Lote{}).Where("cod_lote = ?", loteID).Take(&lote).Error

	if err != nil {
		return &Lote{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Lote{}, errors.New("Lote Not Found")
	}

	return lote, err
}

func (lote *Lote) UpdateLote(db *gorm.DB, loteID uint64) (*Lote, error) {

	db = db.Debug().Model(&Lote{}).Where("cod_lote = ?", loteID).Take(&lote).UpdateColumns(
		map[string]interface{}{
			"cnpj":          lote.Cnpj,
			"contrato":      lote.Contrato,
			"dt_inicio_vig": lote.Dt_inicio_vig,
			"dt_final_vig":  lote.Dt_final_vig,
			"dt_reajuste":   lote.Dt_reajuste,
		},
	)

	if db.Error != nil {
		return &Lote{}, db.Error
	}

	err := db.Debug().Model(&Lote{}).Where("cod_lote = ?", loteID).Take(&lote).Error
	if err != nil {
		return &Lote{}, err
	}

	return lote, err
}
