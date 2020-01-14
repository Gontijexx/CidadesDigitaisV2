package models

import "github.com/jinzhu/gorm"

func (entidade *Entidade) SaveEntidade(db *gorm.DB) (*Entidade, error) {

	err := db.Debug().Create(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}
	return entidade, nil

}
