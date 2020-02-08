package models

import "github.com/jinzhu/gorm"

/*	=========================
		PRECISA FAZER OS TESTES
=========================	*/

/*  =========================
	FUNCAO LISTAR ITENS EMPENHO POR ID
=========================  */

func (itensEmpenho *ItensEmpenho) FindItensEmpenhoByID(db *gorm.DB, idEmpenho, codItem, codTipoItem uint64) (*ItensEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Take(&itensEmpenho).Error

	if err != nil {
		return &ItensEmpenho{}, err
	}

	return itensEmpenho, err
}

/*  =========================
	FUNCAO EDITAR ITENS EMPENHO
=========================  */

func (itensEmpenho *ItensEmpenho) UpdateItensEmpenho(db *gorm.DB, idEmpenho, codItem, codTipoItem uint64) (*ItensEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Updates(
		ItensEmpenho{
			Valor:      itensEmpenho.Valor,
			Quantidade: itensEmpenho.Quantidade}).Error

	if db.Error != nil {
		return &ItensEmpenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Take(&itensEmpenho).Error
	if err != nil {
		return &ItensEmpenho{}, err
	}

	// retorna o elemento que foi alterado
	return itensEmpenho, err
}
