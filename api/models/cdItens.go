package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		PRECISA FAZER OS TESTES
=========================	*/

/*  =========================
	FUNCAO LISTAR CD_ITENS POR ID
=========================  */

func (cdItens *CDItens) FindCDItensByID(db *gorm.DB, cdCodIbge, cdCodItem, cdCodTipoItem uint64) (*CDItens, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdCodIbge, cdCodItem, cdCodTipoItem).Take(&cdItens).Error

	if err != nil {
		return &CDItens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CDItens{}, errors.New("Cd_itens Not Found")
	}

	return cdItens, err
}

/*  =========================
	FUNCAO LISTAR TODOS CD_ITENS POR ID
=========================  */

func (cdItens *CDItens) FindAllCDItens(db *gorm.DB) (*[]CDItens, error) {

	allCDItens := []CDItens{}

	//	Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&CDItens{}).Find(&allCDItens).Error
	if err != nil {
		return &[]CDItens{}, err
	}
	return &allCDItens, err
}

/*  =========================
	FUNCAO EDITAR CD_ITENS POR ID
=========================  */

func (cdItens *CDItens) UpdateCDItens(db *gorm.DB, cdCodIbge, cdCodItem, cdCodTipoItem uint64) (*CDItens, error) {

	err := db.Debug().Model(&CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdCodIbge, cdCodItem, cdCodTipoItem).Updates(
		CDItens{
			QuantidadePrevisto:         cdItens.QuantidadePrevisto,
			QuantidadeProjetoExecutivo: cdItens.QuantidadeProjetoExecutivo,
			QuantidadeTermoInstalacao:  cdItens.QuantidadeTermoInstalacao,
		}).Error

	if db.Error != nil {
		return &CDItens{}, db.Error
	}

	err = db.Debug().Model(&CDItens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdCodIbge, cdCodItem, cdCodTipoItem).Take(&cdItens).Error
	if err != nil {
		return &CDItens{}, err
	}

	return cdItens, err
}
