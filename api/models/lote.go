package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCOES QUE FAZEM CONEXAO DIRETA COM O BANCO DE DADOS
=========================  */

/*  =========================
	FUNCAO PREPARE LOTE
=========================  */

//	Prepara as informacoes a serem escritas no banco de dados
func (lote *Lote) Prepare() {
	lote.Cod_lote = 0
	lote.Cnpj = 0
	lote.Contrato = html.EscapeString(strings.TrimSpace(lote.Contrato))
	lote.Dt_inicio_vig = html.EscapeString(strings.TrimSpace(lote.Dt_inicio_vig))
	lote.Dt_final_vig = html.EscapeString(strings.TrimSpace(lote.Dt_final_vig))
	lote.Dt_reajuste = html.EscapeString(strings.TrimSpace(lote.Dt_reajuste))

}

/*  =========================
	FUNCAO SALVAR LOTE NO BANCO DE DADOS
=========================  */

func (lote *Lote) SaveLote(db *gorm.DB) (*Lote, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&lote).Error
	if err != nil {
		return &Lote{}, err
	}
	return lote, nil

}

/*  =========================
	FUNCAO LISTAR LOTE POR ID
=========================  */

func (lote *Lote) FindLoteByID(db *gorm.DB, loteID uint64) (*Lote, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Lote{}).Where("cod_lote = ?", loteID).Take(&lote).Error

	if err != nil {
		return &Lote{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Lote{}, errors.New("Lote Not Found")
	}

	return lote, err
}

/*  =========================
	FUNCAO LISTAR LOTES
=========================  */

func (lote *Lote) FindLotes(db *gorm.DB) (*[]Lote, error) {

	lotes := []Lote{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Lote{}).Limit(100).Find(&lotes).Error
	if err != nil {
		return &[]Lote{}, err
	}
	return &lotes, err
}

/*  =========================
	FUNCAO EDITAR LOTE
=========================  */

func (lote *Lote) UpdateLote(db *gorm.DB, loteID uint64) (*Lote, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Lote{}).Where("cod_lote = ?", loteID).Take(&lote).UpdateColumns(
		map[string]interface{}{
			"cnpj":          lote.Cnpj,
			"contrato":      lote.Contrato,
			"dt_inicio_vig": lote.Dt_inicio_vig,
			"dt_final_vig":  lote.Dt_final_vig,
			"dt_reajuste":   lote.Dt_reajuste,
		},
	)

	if db.Error != nil {
		return &Lote{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Lote{}).Where("cod_lote = ?", loteID).Take(&lote).Error
	if err != nil {
		return &Lote{}, err
	}

	// retorna o elemento que foi alterado
	return lote, err
}

/*  =========================
	FUNCAO DELETAR LOTE POR ID
=========================  */

func (lote *Lote) DeleteLote(db *gorm.DB, loteID uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Lote{}).Where("cnpj = ?", loteID).Take(&Lote{}).Delete(&Lote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Entidade not found")
		}
		return 0, db.Error
	}

	//	Retornar o numero de linhas deletedas no banco de dados
	return db.RowsAffected, nil
}
