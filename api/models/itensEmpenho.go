package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	STRUCT ITENS EMPENHO
=========================  */

type ItensEmpenho struct {
	IDEmpenho          uint32  `gorm:"primary_key;foreign_key:IDEmpenho;not null" json:"id_empenho"`
	CodItem            uint32  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem        uint32  `gorm:"primary_key;foreign_key:CodTipoItem;not null" json:"cod_tipo_item"`
	CodPrevisaoEmpenho uint32  `gorm:"foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho"`
	Valor              float32 `gorm:"default:null" json:"valor"`
	Quantidade         float32  `gorm:"default:null" json:"quantidade"`
}

/*  =========================
	FUNCAO SALVAR ITENS EMPENHO
=========================  */

func (itensEmpenho *ItensEmpenho) SaveItensEmpenho(db *gorm.DB) (*ItensEmpenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&itensEmpenho).Error
	if err != nil {
		return &ItensEmpenho{}, err
	}

	return itensEmpenho, err
}

/*  =========================
	FUNCAO LISTAR ITENS EMPENHO POR ID
=========================  */

func (itensEmpenho *ItensEmpenho) FindItensEmpenhoByID(db *gorm.DB, idEmpenho, codItem, codTipoItem uint32) (*ItensEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Take(&itensEmpenho).Error
	if err != nil {
		return &ItensEmpenho{}, err
	}

	return itensEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS ITENS EMPENHO
=========================  */

func (itensEmpenho *ItensEmpenho) FindAllItensEmpenho(db *gorm.DB) (*[]ItensEmpenho, error) {

	allItensEmpenho := []ItensEmpenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Entidade{}).Find(&allItensEmpenho).Error
	if err != nil {
		return &[]ItensEmpenho{}, err
	}

	return &allItensEmpenho, err
}

/*  =========================
	FUNCAO EDITAR ITENS EMPENHO
=========================  */

func (itensEmpenho *ItensEmpenho) UpdateItensEmpenho(db *gorm.DB, idEmpenho, codItem, codTipoItem uint32) (*ItensEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE itens_empenho SET cod_previsao_empenho =?, valor = ?, quantidade = ? WHERE id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", itensEmpenho.CodPrevisaoEmpenho, itensEmpenho.Valor, itensEmpenho.Quantidade, idEmpenho, codItem, codTipoItem)
	if db.Error != nil {
		return &ItensEmpenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Take(&itensEmpenho).Error
	if err != nil {
		return &ItensEmpenho{}, err
	}

	// retorna o elemento que foi alterado
	return itensEmpenho, err
}

/*  =========================
	FUNCAO DELETAR ITENS EMPENHO
=========================  */

func (itensEmpenho *ItensEmpenho) DeleteItensEmpenho(db *gorm.DB, idEmpenho, codItem, codTipoItem uint32) error {

	//	Deleta um elemento contido no banco de dados a de suas chaves primarias
	db = db.Debug().Model(&ItensEmpenho{}).Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Take(&ItensEmpenho{}).Delete(&ItensEmpenho{})

	return db.Error
}
