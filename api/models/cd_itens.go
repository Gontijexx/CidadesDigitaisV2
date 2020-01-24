package models

import (
	"errors"
	
	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO LISTAR CD_ITENS
=========================  */

func (e *Cd_itens) FindCdItens(db *gorm.DB) (*[]Cd_itens, error) {
	cdItens := []Cd_itens{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Cd_itens{}).Find(&cdItens).Error
	if err != nil {
		return &[]Cd_itens{}, err
	}
	return &cdItens, err
}

/*  =========================
FUNCAO LISTAR CD_ITENS POR ID
=========================  */

func (cdItens *Cd_itens) FindCdItem(db *gorm.DB, cdItensID1, cdItensID2, cdItensID3 uint64) (*Cd_itens, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Cd_itens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID1, cdItensID2, cdItensID3).Take(&cdItens).Error

	if err != nil {
		return &Cd_itens{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cd_itens{}, errors.New("Cd_itens Not Found")
	}

	return cdItens, err
}

/*  =========================
	FUNCAO ATUALIZAR CD_ITENS
=========================  */

func (cdItens *Cd_itens) UpdateCdItem(db *gorm.DB, cdItensID1, cdItensID2, cdItensID3 uint64) (*Cd_itens, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Cd_itens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID1, cdItensID2, cdItensID3).Take(&cdItens).UpdateColumns(
		map[string]interface{}{
			"quantidade_previsto":          cdItens.Quantidade_previsto,
			"quantidade_projeto_executivo": cdItens.Quantidade_projeto_executivo,
			"quantidade_termo_instalacao":  cdItens.Quantidade_termo_instalacao,
		},
	)

	if db.Error != nil {
		return &Cd_itens{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Cd_itens{}).Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item =?", cdItensID1, cdItensID2, cdItensID3).Take(&cdItens).Error
	if err != nil {
		return &Cd_itens{}, err
	}

	// retorna o elemento que foi alterado
	return cdItens, err
}
