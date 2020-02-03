package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO LISTAR ITENS PREVISAO EMPENHO POR ID
=========================  */

func (itensPrevisaoEmpenho *ItensPrevisaoEmpenho) FindItensPrevisaoEmpenhoByID(db *gorm.DB, itensPrevisaoEmpenhoID uint64) (*ItensPrevisaoEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(ItensPrevisaoEmpenho{}).Where("cod_previsao_empenho", itensPrevisaoEmpenhoID).Take(&itensPrevisaoEmpenho).Error

	if err != nil {
		return &ItensPrevisaoEmpenho{}, err
	}

	return itensPrevisaoEmpenho, err
}

/*  =========================
	FUNCAO EDITAR ITENS PREVISAO EMPENHO
=========================  */

func (itensPrevisaoEmpenho *ItensPrevisaoEmpenho) UpdateItensPrevisaoEmpenho(db *gorm.DB, itensPrevisaoEmpenhoID uint64) (*ItensPrevisaoEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&ItensPrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", itensPrevisaoEmpenhoID).Updates(
		ItensPrevisaoEmpenho{
			Valor:      itensPrevisaoEmpenho.Valor,
			Quantidade: itensPrevisaoEmpenho.Quantidade,
		}).Error

	if db.Error != nil {
		return &ItensPrevisaoEmpenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&ItensPrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", itensPrevisaoEmpenhoID).Take(&itensPrevisaoEmpenho).Error
	if err != nil {
		return &ItensPrevisaoEmpenho{}, err
	}

	// retorna o elemento que foi alterado
	return itensPrevisaoEmpenho, err
}
