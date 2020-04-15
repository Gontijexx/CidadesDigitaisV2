package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ITENS FATURA NO BANCO DE DADOS
=========================  */

func (itensFatura *ItensFatura) SaveItensFatura(db *gorm.DB) (*ItensFatura, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&itensFatura).Error
	if err != nil {
		return &ItensFatura{}, err
	}

	return itensFatura, nil
}

/*  =========================
	FUNCAO LISTAR ITENS FATURA POR ID
=========================  */

func (itensFatura *ItensFatura) FindItensFaturaByID(db *gorm.DB, numNF, codIbge, idEmpenho, codItem, codTipoItem uint64) (*ItensFatura, error) {

	//	Busca um elemento no banco de dados de acordo com suas chaves primarias
	err := db.Debug().Model(ItensEmpenho{}).Where("num_nf = ? AND cod_ibge = ? AND id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", numNF, codIbge, idEmpenho, codItem, codTipoItem).Take(&itensFatura).Error

	if err != nil {
		return &ItensFatura{}, err
	}

	return itensFatura, err
}

/*  =========================
	FUNCAO LISTAR TODAS ITENS FATURA
=========================  */

func (itensFatura *ItensFatura) FindAllItensFatura(db *gorm.DB) (*[]ItensFatura, error) {

	allItensFatura := []ItensFatura{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&ItensFatura{}).Find(&allItensFatura).Error
	if err != nil {
		return &[]ItensFatura{}, err
	}
	return &allItensFatura, err
}

/*  =========================
	FUNCAO EDITAR ITENS FATURA
=========================  */

func (itensFatura *ItensFatura) UpdateItensFatura(db *gorm.DB, numNF, codIbge, idEmpenho, codItem, codTipoItem uint64) (*ItensFatura, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE itens_fatura SET valor = ?, quantidade = ? WHERE num_nf = ? AND cod_ibge = ? AND id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", itensFatura.Valor, itensFatura.Quantidade, numNF, codIbge, idEmpenho, codItem, codTipoItem)
	if db.Error != nil {
		return &ItensFatura{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&ItensFatura{}).Take(&itensFatura).Error
	if err != nil {
		return &ItensFatura{}, err
	}

	// retorna o elemento que foi alterado
	return itensFatura, err
}

/*  =========================
	FUNCAO DELETAR ITENS OTB POR ID
=========================  */

func (itensFatura *ItensFatura) DeleteItensFatura(db *gorm.DB, numNF, cod_ibge, idEmpenho, codItem, codTipoItem uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&ItensFatura{}).Where("num_nf = ? AND cod_ibge = ? AND id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", numNF, cod_ibge, idEmpenho, codItem, codTipoItem).Take(&ItensFatura{}).Delete(&ItensFatura{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ItensFatura not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
