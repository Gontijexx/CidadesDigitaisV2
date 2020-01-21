package models

import (
	"errors"

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

func (cd *Cd) UpdateCd(db *gorm.DB, cdID uint64) (*Cd, error) {

	db = db.Debug().Model(&Cd{}).Where("cod_ibge = ?", cdID).Take(&cd).UpdateColumns(
		map[string]interface{}{
			"cod_lote": cd.Cod_lote,
			"os_pe":    cd.Os_pe,
			"data_pe":  cd.Data_pe,
			"os_imp":   cd.Os_imp,
			"data_imp": cd.Data_imp,
		},
	)

	if db.Error != nil {
		return &Cd{}, db.Error
	}

	err := db.Debug().Model(&Cd{}).Where("cod_ibge = ?", cdID).Take(&cd).Error
	if err != nil {
		return &Cd{}, err
	}

	return cd, err
}
