package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

func (cd *Cd) SaveCd(db *gorm.DB) (*Cd, error) {

	err := db.Debug().Create(&cd).Error
	if err != nil {
		return &Cd{}, err
	}
	return cd, nil

}

func (cd *Cd) FindCdByID(db *gorm.DB, cdID uint64) (*Cd, error) {

	err := db.Debug().Model(Cd{}).Where("cod_ibge = ?", cdID).Take(&cd).Error

	if err != nil {
		return &Cd{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cd{}, errors.New("Cd Not Found")
	}

	return cd, err
}

func (c *Cd) UpdateACd(db *gorm.DB, cId uint32) (*Cd, error) {

	// To hash the password
	err := c.BeforeSave()
	if err != nil {
		log.Printf("[FATAL] cannot HASH password, %v\n", err)
	}

	db = db.Debug().Model(&Cd{}).Where("cod_ibge= ?", cId).Take(&Cd{}).UpdateColumns(
		map[string]interface{}{
			"cod_lote": c.Cod_lote,
			"os_pe":    c.Os_pe,
			"data_pe":  c.Data_pe,
			"os_imp":   c.Os_imp,
			"data_imp": c.Data_imp,
		},
	)
	if db.Error != nil {
		return &Cd{}, db.Error
	}

	// This is the display the updated user
	err = db.Debug().Model(&Cd{}).Where("cod_ibge = ?", cId).Take(&c).Error
	if err != nil {
		return &Cd{}, err
	}

	return c, nil

}
