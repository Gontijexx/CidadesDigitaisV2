package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PREVISAO
	EMPENHO	NO BANCO DE DADOS
=========================  */

func (previsaoEmpenho *Previsao_empenho) SavePrevisaoEmpenho(db *gorm.DB) (*Previsao_empenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&previsaoEmpenho).Error
	if err != nil {
		return &Previsao_empenho{}, err
	}
	return previsaoEmpenho, nil

}

/*  =========================
	FUNCAO LISTAR
	PREVISAO EMPENHO POR ID
=========================  */

func (previsaoEmpenho *Previsao_empenho) FindPrevisaoEmpenhoByID(db *gorm.DB, previsaoEmpenhoID uint64) (*Previsao_empenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Previsao_empenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error

	if err != nil {
		return &Previsao_empenho{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Previsao_empenho{}, errors.New("Previsao_empenho Not Found")
	}

	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *Previsao_empenho) FindAllPrevisaoEmpenho(db *gorm.DB) (*[]Previsao_empenho, error) {

	previsaoEmpenhos := []Previsao_empenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Previsao_empenho{}).Find(&previsaoEmpenhos).Error
	if err != nil {
		return &[]Previsao_empenho{}, err
	}
	return &previsaoEmpenhos, err
}

/*  =========================
	FUNCAO EDITAR PREVISAO_EMPENHO
=========================  */

func (previsao_empenho *Previsao_empenho) UpdatePrevisaoEmpenho(db *gorm.DB, previsao_empenhoID uint64) (*Previsao_empenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Previsao_empenho{}).Where("cod_previsao_empenho = ?", previsao_empenhoID).Take(&Previsao_empenho{}).UpdateColumns(
		map[string]interface{}{
			"data":         			previsao_empenho.Data,
			"tipo":         			previsao_empenho.Tipo,
			"ano_referencia":           previsao_empenho.Ano_referencia,
		},
	)

	if db.Error != nil {
		return &Previsao_empenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Previsao_empenho{}).Where("cod_previsao_empenho = ?", previsao_empenhoID).Take(&previsao_empenho).Error
	if err != nil {
		return &Previsao_empenho{}, err
	}

	// retorna o elemento que foi alterado
	return previsao_empenho, err
}
