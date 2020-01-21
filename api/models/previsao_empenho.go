package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func (previsaoEmpenho *Previsao_empenho) SavePrevisao_empenho(db *gorm.DB) (*Previsao_empenho, error) {

	err := db.Debug().Create(&previsaoEmpenho).Error
	if err != nil {
		return &Previsao_empenho{}, err
	}
	return previsaoEmpenho, nil

}

func (previsaoEmpenho *Previsao_empenho) FindPrevisao_empenhoByID(db *gorm.DB, previsaoEmpenhoID uint64) (*Previsao_empenho, error) {

	err := db.Debug().Model(Previsao_empenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error

	if err != nil {
		return &Previsao_empenho{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Previsao_empenho{}, errors.New("Previsao_empenho Not Found")
	}

	return previsaoEmpenho, err
}

func (previsaoEmpenho *Previsao_empenho) UpdatePrevisao_empenho(db *gorm.DB, previsaoEmpenhoID uint64) (*Previsao_empenho, error) {

	db = db.Debug().Model(&Previsao_empenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).UpdateColumns(
		map[string]interface{}{
			"cod_lote":             previsaoEmpenho.Cod_lote,
			"cod_natureza_despesa": previsaoEmpenho.Cod_natureza_despesa,
			"data":                 previsaoEmpenho.Data,
			"tipo":                 previsaoEmpenho.Tipo,
			"ano_referencia":       previsaoEmpenho.Ano_referencia,
		},
	)

	if db.Error != nil {
		return &Previsao_empenho{}, db.Error
	}

	err := db.Debug().Model(&Previsao_empenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error
	if err != nil {
		return &Previsao_empenho{}, err
	}

	return previsaoEmpenho, err
}

func (e *Previsao_empenho) FindAllPrevisao_empenho(db *gorm.DB) (*[]Previsao_empenho, error) {
	previsaoEmpenho := []Previsao_empenho{}
	err := db.Debug().Model(&Previsao_empenho{}).Find(&previsaoEmpenho).Error
	if err != nil {
		return &[]Previsao_empenho{}, err
	}
	return &previsaoEmpenho, err
}
