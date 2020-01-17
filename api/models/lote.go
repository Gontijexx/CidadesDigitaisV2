package models

import (
	"errors"
	"log"

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

func (l *Lote) UpdateALote(db *gorm.DB, lId uint32) (*Lote, error) {

	// To hash the password
	err := l.BeforeSave()
	if err != nil {
		log.Printf("[FATAL] cannot HASH password, %v\n", err)
	}

	db = db.Debug().Model(&Lote{}).Where("cod_lote= ?", lId).Take(&Lote{}).UpdateColumns(
		map[string]interface{}{
			"cnpj":          l.Cnpj,
			"contrato":      l.Contrato,
			"dt_inicio_vig": l.Dt_inicio_vig,
			"dt_final_vig":  l.Dt_final_vig,
			"dt_reajuste":   l.Dt_reajuste,
		},
	)
	if db.Error != nil {
		return &Lote{}, db.Error
	}

	// This is the display the updated user
	err = db.Debug().Model(&Lote{}).Where("cod_lote = ?", lId).Take(&l).Error
	if err != nil {
		return &Lote{}, err
	}

	return l, nil

}
