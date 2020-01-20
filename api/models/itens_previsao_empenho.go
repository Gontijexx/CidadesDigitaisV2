package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func (itensPrevisaoEmpenho *Itens_previsao_empenho) FindItens_previsao_empenhoByID(db *gorm.DB, itensPrevisaoEmpenhoID1, itensPrevisaoEmpenhoID2, itensPrevisaoEmpenhoID3 uint64) (*Itens_previsao_empenho, error) {

	err := db.Debug().Model(Itens_previsao_empenho{}).Where("cod_previsao_empenho = ? AND cod_item = ? AND cod_tipo_item =?", itensPrevisaoEmpenhoID1, itensPrevisaoEmpenhoID2, itensPrevisaoEmpenhoID3).Take(&itensPrevisaoEmpenho).Error

	if err != nil {
		return &Itens_previsao_empenho{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Itens_previsao_empenho{}, errors.New("Itens_previsao_empenho Not Found")
	}

	return itensPrevisaoEmpenho, err
}

func (itensPrevisaoEmpenho *Itens_previsao_empenho) UpdateItens_previsao_empenho(db *gorm.DB, itensPrevisaoEmpenhoID1, itensPrevisaoEmpenhoID2, itensPrevisaoEmpenhoID3 uint64) (*Itens_previsao_empenho, error) {

	db = db.Debug().Model(&Itens_previsao_empenho{}).Where("cod_previsao_empenho = ? AND cod_item = ? AND cod_tipo_item =?", itensPrevisaoEmpenhoID1, itensPrevisaoEmpenhoID2, itensPrevisaoEmpenhoID3).Take(&itensPrevisaoEmpenho).UpdateColumns(
		map[string]interface{}{
			"cod_lote":   itensPrevisaoEmpenho.Cod_lote,
			"valor":      itensPrevisaoEmpenho.Valor,
			"quantidade": itensPrevisaoEmpenho.Quantidade,
		},
	)

	if db.Error != nil {
		return &Itens_previsao_empenho{}, db.Error
	}

	err := db.Debug().Model(&Itens_previsao_empenho{}).Where("cod_previsao_empenho = ? AND cod_item = ? AND cod_tipo_item =?", itensPrevisaoEmpenhoID1, itensPrevisaoEmpenhoID2, itensPrevisaoEmpenhoID3).Take(&itensPrevisaoEmpenho).Error
	if err != nil {
		return &Itens_previsao_empenho{}, err
	}

	return itensPrevisaoEmpenho, err
}

func (e *Itens_previsao_empenho) FindAllItens_previsao_empenho(db *gorm.DB) (*[]Itens_previsao_empenho, error) {
	itensPrevisaoEmpenho := []Itens_previsao_empenho{}
	err := db.Debug().Model(&Itens_previsao_empenho{}).Find(&itensPrevisaoEmpenho).Error
	if err != nil {
		return &[]Itens_previsao_empenho{}, err
	}
	return &itensPrevisaoEmpenho, err
}
