package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO LISTAR LOTE_ITENS POR ID
=========================  */

func (loteItens *LoteItens) FindLoteItensByID(db *gorm.DB, loteCodLote, loteCodItem, loteCodTipoItem uint64) (*LoteItens, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteCodLote, loteCodItem, loteCodTipoItem).Take(&loteItens).Error

	if err != nil {
		return &LoteItens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &LoteItens{}, errors.New("Lote_itens Not Found")
	}

	return loteItens, err
}

/*  =========================
	FUNCAO LISTAR TODOS LOTE ITENS
=========================  */

func (loteItens *LoteItens) FindAllLoteItens(db *gorm.DB) (*[]LoteItens, error) {

	allLoteItens := []LoteItens{}

	//	Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&LoteItens{}).Find(&allLoteItens).Error
	if err != nil {
		return &[]LoteItens{}, err
	}
	return &allLoteItens, err
}

/*  =========================
	FUNCAO EDITAR LOTE_ITENS
=========================  */

func (loteItens *LoteItens) UpdateLoteItens(db *gorm.DB, loteCodLote, loteCodItem, loteCodTipoItem uint64) (*LoteItens, error) {

	err := db.Debug().Model(&LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteCodLote, loteCodItem, loteCodTipoItem).Updates(
		LoteItens{
			Preco: loteItens.Preco,
		}).Error

	if db.Error != nil {
		return &LoteItens{}, db.Error
	}

	err = db.Debug().Model(&LoteItens{}).Where("cod_lote = ? AND cod_item = ? AND cod_tipo_item =?", loteCodLote, loteCodItem, loteCodTipoItem).Take(&loteItens).Error
	if err != nil {
		return &LoteItens{}, err
	}

	return loteItens, err
}
