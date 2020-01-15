package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func (entidade *Entidade) SaveEntidade(db *gorm.DB) (*Entidade, error) {

	err := db.Debug().Create(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}
	return entidade, nil

}

func (entidade *Entidade) FindEntidadeByID(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	err := db.Debug().Model(Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error

	if err != nil {
		return &Entidade{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Entidade{}, errors.New("Entidade Not Found")
	}

	return entidade, err
}
