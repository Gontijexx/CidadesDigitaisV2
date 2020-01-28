package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		COMENTAR
=========================	*/

func (reajuste *Reajuste) SaveReajuste(db *gorm.DB) (*Reajuste, error) {

	err := db.Debug().Create(&reajuste).Error
	if err != nil {
		return &Reajuste{}, err
	}
	return reajuste, nil

}

func (reajuste *Reajuste) FindReajusteByID(db *gorm.DB, reajusteID uint64) (*Reajuste, error) {

	err := db.Debug().Model(Reajuste{}).Where("ano_ref = ?", reajusteID).Take(&reajuste).Error

	if err != nil {
		return &Reajuste{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Reajuste{}, errors.New("Reajuste Not Found")
	}

	return reajuste, err
}

func (r *Reajuste) UpdateReajuste(db *gorm.DB, rId1, rId2 uint32) (*Reajuste, error) {

	db = db.Debug().Model(&Reajuste{}).Where("ano_ref= ? AND cod_lote= ?", rId1, rId2).Take(&Reajuste{}).UpdateColumns(
		map[string]interface{}{
			"percentual": r.Percentual,
		},
	)

	if db.Error != nil {
		return &Reajuste{}, db.Error
	}

	err := db.Debug().Model(&Reajuste{}).Where("ano_ref= ? AND cod_lote= ?", rId1, rId2).Take(&r).Error
	if err != nil {
		return &Reajuste{}, err
	}

	return r, err

}

func (r *Reajuste) DeleteReajuste(db *gorm.DB, rId1 uint32, rId2 int32) (int64, error) {

	db = db.Debug().Model(&Reajuste{}).Where("ano_ref = ? AND cod_lote", rId1, rId2).Take(&Reajuste{}).Delete(&Reajuste{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
