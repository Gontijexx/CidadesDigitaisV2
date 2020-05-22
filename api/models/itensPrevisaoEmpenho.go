package models

import "github.com/jinzhu/gorm"

/*  =========================
	STRUTC ITENS PREVISAO EMPENHO
=========================  */

type ItensPrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint32  `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho"`
	CodItem            uint32  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem        uint32  `gorm:"primary_key;foreign_key:CodTipo_item;not null" json:"cod_tipo_item"`
	CodLote            uint32  `gorm:"foreign_key:CodLote;not null" json:"cod_lote"`
	Valor              float32 `gorm:"default:null" json:"valor"`
	Quantidade         float32 `gorm:"default:null" json:"quantidade"`
	Descricao          string  `gorm:"default:null" json:"descricao"`
}

/*  =========================
	FUNCAO LISTAR ITENS PREVISAO EMPENHO POR ID
=========================  */

func (itensPrevisaoEmpenho *ItensPrevisaoEmpenho) FindItensPrevisaoEmpenhoByID(db *gorm.DB, codPrevisaoEmpenho, codItem, codTipoItem uint32) (*ItensPrevisaoEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&ItensPrevisaoEmpenho{}).Where("cod_previsao_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", codPrevisaoEmpenho, codItem, codTipoItem).Take(&itensPrevisaoEmpenho).Error

	if err != nil {
		return &ItensPrevisaoEmpenho{}, err
	}

	return itensPrevisaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODOS ITENS PREVISAO EMPENHO
=========================  */

func (itensPrevisaoEmpenho *ItensPrevisaoEmpenho) FindAllItensPrevisaoEmpenho(db *gorm.DB) (*[]ItensPrevisaoEmpenho, error) {

	allItensPrevisaoEmpenho := []ItensPrevisaoEmpenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Table("itens_previsao_empenho").
		Select("itens.descricao, itens_previsao_empenho.*").
		Joins("JOIN itens ON itens_previsao_empenho.cod_item = itens.cod_item AND itens_previsao_empenho.cod_tipo_item = itens.cod_tipo_item").
		Scan(&allItensPrevisaoEmpenho).Error
	if err != nil {
		return &[]ItensPrevisaoEmpenho{}, err
	}

	return &allItensPrevisaoEmpenho, err
}

/*  =========================
	FUNCAO EDITAR ITENS PREVISAO EMPENHO
=========================  */

func (itensPrevisaoEmpenho *ItensPrevisaoEmpenho) UpdateItensPrevisaoEmpenho(db *gorm.DB, codPrevisaoEmpenho, codItem, codTipoItem uint64) (*ItensPrevisaoEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE itens_previsao_empenho SET valor = ?, quantidade = ? WHERE cod_previsao_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", itensPrevisaoEmpenho.Valor, itensPrevisaoEmpenho.Quantidade, codPrevisaoEmpenho, codItem, codTipoItem)
	if db.Error != nil {
		return &ItensPrevisaoEmpenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&ItensPrevisaoEmpenho{}).Where("cod_previsao_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", codPrevisaoEmpenho, codItem, codTipoItem).Take(&itensPrevisaoEmpenho).Error
	if err != nil {
		return &ItensPrevisaoEmpenho{}, err
	}

	return itensPrevisaoEmpenho, err
}
