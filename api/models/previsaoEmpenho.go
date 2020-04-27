package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	STRUCT PREVISAO EMPENHO
=========================  */

type PrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint32 `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;auto_increment;not null" json:"cod_previsao_empenho"`
	CodLote            uint32 `gorm:"foreign_key:CodLote;not null" json:"cod_lote"`
	CodNaturezaDespesa uint32 `gorm:"foreign_key:CodNaturezaDespesa;not null" json:"cod_natureza_despesa"`
	Data               string `gorm:"default:null" json:"data"`
	Tipo               string `gorm:"default:null" json:"tipo"`
	Ano_referencia     uint32 `gorm:"default:null" json:"ano_referencia"`
}

/*  =========================
	FUNCAO SALVAR PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) SavePrevisaoEmpenho(db *gorm.DB) (*PrevisaoEmpenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}

	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR PREVISAO EMPENHO POR ID
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) FindPrevisaoEmpenhoByID(db *gorm.DB, codPrevisaoEmpenho uint32) (*PrevisaoEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", codPrevisaoEmpenho).Take(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}

	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) FindAllPrevisaoEmpenho(db *gorm.DB) (*[]PrevisaoEmpenho, error) {

	allPrevisaoEmpenho := []PrevisaoEmpenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Table("previsao_empenho").Select("natureza_despesa.descricao, previsao_empenho.*").
		Joins("JOIN natureza_despesa ON previsao_empenho.cod_natureza_despesa = natureza_despesa.cod_natureza_despesa ORDER BY previsao_empenho.cod_previsao_empenho ASC").Scan(&allPrevisaoEmpenho).Error
	if err != nil {
		return &[]PrevisaoEmpenho{}, err
	}

	return &allPrevisaoEmpenho, err
}

/*  =========================
	FUNCAO EDITAR PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) UpdatePrevisaoEmpenho(db *gorm.DB, codPrevisaoEmpenho uint32) (*PrevisaoEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE previsao_empenho SET cod_lote = ?, cod_natureza_despesa = ? ,data = ?, tipo = ?, ano_referencia = ? WHERE cod_previsao_empenho = ?", previsaoEmpenho.CodLote, previsaoEmpenho.CodNaturezaDespesa, previsaoEmpenho.Data, previsaoEmpenho.Tipo, previsaoEmpenho.Ano_referencia, codPrevisaoEmpenho)
	if db.Error != nil {
		return &PrevisaoEmpenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", codPrevisaoEmpenho).Take(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}

	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO DELETAR PREVISAO EMPENHO
=========================

func (previsaoEmpenho *PrevisaoEmpenho) DeletePrevisaoEmpenho(db *gorm.DB, codPrevisaoEmpenho uint32) error {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", codPrevisaoEmpenho).Take(&PrevisaoEmpenho{}).Delete(&PrevisaoEmpenho{})

	return db.Error
}

*/
