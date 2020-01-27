package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO LISTAR ITENS PREVISAO EMPENHO POR ID
=========================  */


func (itensPrevisaoEmpenho *Itens_previsao_empenho) FindItensPrevisaoEmpenho(db *gorm.DB, itensPrevisaoEmpenhoID uint64) (*Itens_previsao_empenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Itens_previsao_empenho{}).Where("cod_previsao_empenho", itensPrevisaoEmpenhoID).Take(&itensPrevisaoEmpenho).Error

	if err != nil {
		return &Itens_previsao_empenho{}, err
	}

	return itensPrevisaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODOS ITENS PREVISAO EMPENHO
=========================  */


func (itensPrevisaoEmpenho *Itens_previsao_empenho) FindAllItensPrevisaoEmpenho(db *gorm.DB) (*[]Itens_previsao_empenho, error) {

	itensPrevisaoEmpenhos := []Itens_previsao_empenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Itens_previsao_empenho{}).Limit(100).Find(&itensPrevisaoEmpenhos).Error
	if err != nil {
		return &[]Itens_previsao_empenho{}, err
	}
	return &itensPrevisaoEmpenhos, err
}

/*  =========================
	FUNCAO EDITAR ITENS PREVISAO EMPENHO
=========================  */


func (itensPrevisaoEmpenho *Itens_previsao_empenho) UpdateItensPrevisaoEmpenho(db *gorm.DB, itensPrevisaoEmpenhoID uint64) (*Itens_previsao_empenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Itens_previsao_empenho{}).Where("cod_previsao_empenho = ?", itensPrevisaoEmpenhoID).Take(&itensPrevisaoEmpenho).UpdateColumns(
    
		map[string]interface{}{
			"valor":      itensPrevisaoEmpenho.Valor,
			"quantidade": itensPrevisaoEmpenho.Quantidade,
		},
	)

	if db.Error != nil {
		return &Itens_previsao_empenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Itens_previsao_empenho{}).Where("cod_previsao_empenho = ?", itensPrevisaoEmpenhoID).Take(&itensPrevisaoEmpenho).Error
	if err != nil {
		return &Itens_previsao_empenho{}, err
	}

	// retorna o elemento que foi alterado
	return itensPrevisaoEmpenho, err
}
