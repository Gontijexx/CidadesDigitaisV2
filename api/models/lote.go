package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

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

func (lote *Lote) FindLoteByID(db *gorm.DB, codLote uint64) (*Lote, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Lote{}).Where("cod_lote = ?", codLote).Take(&lote).Error

	if err != nil {
		return &Lote{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Lote{}, errors.New("Lote Not Found")
	}

	return lote, err
}

/*  =========================
	FUNCAO LISTAR TODOS LOTE
=========================  */

func (lote *Lote) FindAllLote(db *gorm.DB) (*[]Lote, error) {

	allLote := []Lote{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Lote{}).Find(&allLote).Error
	if err != nil {
		return &[]Lote{}, err
	}
	return &allLote, err
}

/*  =========================
	FUNCAO EDITAR LOTE
=========================  */

func (lote *Lote) UpdateLote(db *gorm.DB, codLote uint64) (*Lote, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Lote{}).Where("cod_lote = ?", codLote).Updates(
		Lote{
			Cnpj:        lote.Cnpj,
			Contrato:    lote.Contrato,
			DtInicioVig: lote.DtInicioVig,
			DtFinalVig:  lote.DtFinalVig,
			DtReajuste:  lote.DtReajuste}).Error

	if err != nil {
		return &Lote{}, err
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Lote{}).Where("cod_lote = ?", codLote).Take(&lote).Error
	if err != nil {
		return &Lote{}, err
	}

	// retorna o elemento que foi alterado
	return lote, err
}

//Não é possivel deletar por causa das triggers da lote_itens

/*  =========================
	FUNCAO DELETAR LOTE POR ID
=========================

func (lote *Lote) DeleteLote(db *gorm.DB, codLote uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Lote{}).Where("cod_lote = ?", codLote).Take(&Lote{}).Delete(&Lote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Lote not found")
		}
		return 0, db.Error
	}

	//	Retornar o numero de linhas deletadas no banco de dados
	return db.RowsAffected, nil
}
*/
