package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ITENS EMPENHO NO BANCO DE DADOS
=========================  */

func (itensEmpenho *ItensEmpenho) SaveItensEmpenho(db *gorm.DB) (*ItensEmpenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&itensEmpenho).Error
	if err != nil {
		return &ItensEmpenho{}, err
	}

	return itensEmpenho, nil
}

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
	err := db.Debug().Exec("UPDATE itens_empenho SET cod_previsao_empenho =?, valor = ?, quantidade = ? WHERE id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", itensEmpenho.CodPrevisaoEmpenho, itensEmpenho.Valor, itensEmpenho.Quantidade, idEmpenho, codItem, codTipoItem).Error

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

/*  =========================
	FUNCAO DELETAR ITENS EMPENHO
=========================  */

func (itensEmpenho *ItensEmpenho) DeleteItensEmpenho(db *gorm.DB, idEmpenho, codItem, codTipoItem uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a de suas chaves primarias
	db = db.Debug().Model(&ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Take(&ItensEmpenho{}).Delete(&ItensEmpenho{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Itens Empenho not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil

}
