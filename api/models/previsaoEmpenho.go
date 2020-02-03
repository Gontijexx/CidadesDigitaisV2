package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PREVISAO EMPENHO	NO BANCO DE DADOS
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) SavePrevisaoEmpenho(db *gorm.DB) (*PrevisaoEmpenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}
	return previsaoEmpenho, nil

}

/*  =========================
	FUNCAO LISTAR PREVISAO EMPENHO POR ID
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) FindPrevisaoEmpenhoByID(db *gorm.DB, previsaoEmpenhoID uint64) (*PrevisaoEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error

	if err != nil {
		return &PrevisaoEmpenho{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &PrevisaoEmpenho{}, errors.New("Previsao_empenho Not Found")
	}

	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) FindAllPrevisaoEmpenho(db *gorm.DB) (*[]PrevisaoEmpenho, error) {

	allPrevisaoEmpenho := []PrevisaoEmpenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&PrevisaoEmpenho{}).Find(&allPrevisaoEmpenho).Error
	if err != nil {
		return &[]PrevisaoEmpenho{}, err
	}
	return &allPrevisaoEmpenho, err
}

/*  =========================
	FUNCAO EDITAR PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) UpdatePrevisaoEmpenho(db *gorm.DB, previsaoEmpenhoID uint64) (*PrevisaoEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Updates(
		PrevisaoEmpenho{
			Data:           previsaoEmpenho.Data,
			Tipo:           previsaoEmpenho.Tipo,
			Ano_referencia: previsaoEmpenho.Ano_referencia}).Error

	if err != nil {
		return &PrevisaoEmpenho{}, err
	}
	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}

	// retorna o elemento que foi alterado
	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO DELETAR PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) DeletePrevisaoEmpenho(db *gorm.DB, codPrevisaoEmpenho uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", codPrevisaoEmpenho).Take(&PrevisaoEmpenho{}).Delete(&PrevisaoEmpenho{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Previsao_Empenho not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
