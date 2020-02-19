package models

import "github.com/jinzhu/gorm"

/*  =========================
	FUNCAO LISTAR TODAS ITENS_OTB
=========================  */

func (itensOTB *ItensOTB) FindItensOTB(db *gorm.DB, codOTB uint64) (*[]ItensOTB, error) {

	allItensOTB := []ItensOTB{}

	//	Busca todos elementos contidos no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&ItensOTB{}).Where("cod_otb = ?", codOTB).Find(&allItensOTB).Error
	if err != nil {
		return &[]ItensOTB{}, err
	}

	return &allItensOTB, err
}

/*  =========================
	FUNCAO EDITAR ITENS_OTB
=========================  */

func (itensOTB *ItensOTB) UpdateItensOTB(db *gorm.DB) (*ItensOTB, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&ItensOTB{}).Updates(
		ItensOTB{
			Valor:      itensOTB.Valor,
			Quantidade: itensOTB.Quantidade}).Error

	if db.Error != nil {
		return &ItensOTB{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&ItensOTB{}).Take(&itensOTB).Error
	if err != nil {
		return &ItensOTB{}, err
	}

	// retorna o elemento que foi alterado
	return itensOTB, err
}
