package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ITENS OTB NO BANCO DE DADOS
=========================  */

func (itensOTB *ItensOTB) SaveItensOTB(db *gorm.DB) (*ItensOTB, error) {

	//adiciona um novo elemento no banco de dados
	err := db.Debug().Create(&itensOTB).Error
	if err != nil {
		return &ItensOTB{}, err
	}

	return itensOTB, nil
}

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

func (itensOTB *ItensOTB) UpdateItensOTB(db *gorm.DB, codOtb, numNf, codIbge, idEmpenho, codItem, codTipoItem uint64) (*ItensOTB, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE itens_otb SET valor = ?, quantidade = ? WHERE cod_otb = ? AND num_nf = ? AND  cod_ibge = ? AND id_empenho = ? AND cod_item = ?  AND cod_tipo_item = ?", itensOTB.Valor, itensOTB.Quantidade, codOtb, numNf, codIbge, idEmpenho, codItem, codTipoItem).Error

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

/*  =========================
	FUNCAO DELETAR ITENS OTB POR ID
=========================  */

func (itensOTB *ItensOTB) DeleteItensOTB(db *gorm.DB, codOtb, numNf, codIbge, idEmpenho, codItem, codTipoItem uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&ItensOTB{}).Where("cod_otb = ? AND num_nf = ? AND  cod_ibge = ? AND id_empenho = ? AND cod_item = ?  AND cod_tipo_item = ?", codOtb, numNf, codIbge, idEmpenho, codItem, codTipoItem).Take(&ItensOTB{}).Delete(&ItensOTB{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Itens OTB not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
