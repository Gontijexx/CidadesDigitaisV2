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

func (entidade *Entidade) UpdateEntidade(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	db = db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).UpdateColumns(
		map[string]interface{}{
			"nome":           entidade.Nome,
			"endereco":       entidade.Endereco,
			"numero":         entidade.Numero,
			"bairro":         entidade.Bairro,
			"cep":            entidade.Cep,
			"nome_municipio": entidade.Nome_municipio,
			"uf":             entidade.Uf,
			"observacao":     entidade.Observacao,
		},
	)

	if db.Error != nil {
		return &Entidade{}, db.Error
	}

	err := db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}

	return entidade, err
}

func (e *Entidade) FindAllEntidade(db *gorm.DB) (*[]Entidade, error) {
	entidade := []Entidade{}
	err := db.Debug().Model(&Entidade{}).Find(&entidade).Error
	if err != nil {
		return &[]Entidade{}, err
	}
	return &entidade, err
}
