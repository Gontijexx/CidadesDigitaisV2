package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR NATUREZA_DESPESA NO BANCO DE DADOS
=========================  */

func (naturezaDespesa *NaturezaDespesa) SaveNaturezaDespesa(db *gorm.DB) (*NaturezaDespesa, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&naturezaDespesa).Error
	if err != nil {
		return &NaturezaDespesa{}, err
	}
	return naturezaDespesa, nil

}

/*  =========================
	FUNCAO LISTAR NATUREZA_DESPESA POR ID
=========================  */

func (naturezaDespesa *NaturezaDespesa) FindNaturezaDespesaByID(db *gorm.DB, codNaturezaDespesa uint64) (*NaturezaDespesa, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(NaturezaDespesa{}).Where("cod_natureza_despesa = ?", codNaturezaDespesa).Take(&naturezaDespesa).Error

	if err != nil {
		return &NaturezaDespesa{}, err
	}

	return naturezaDespesa, err
}

/*  =========================
	FUNCAO LISTAR TODAS NATUREZA_DESPESA
=========================  */

func (naturezaDespesa *NaturezaDespesa) FindAllNaturezaDespesa(db *gorm.DB) (*[]NaturezaDespesa, error) {

	allNaturezaDespesa := []NaturezaDespesa{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&NaturezaDespesa{}).Find(&allNaturezaDespesa).Error
	if err != nil {
		return &[]NaturezaDespesa{}, err
	}
	return &allNaturezaDespesa, err
}

/*  =========================
	FUNCAO EDITAR NATUREZA_DESPESA
=========================  */

func (naturezaDespesa *NaturezaDespesa) UpdateNaturezaDespesa(db *gorm.DB, codNaturezaDespesa uint64) (*NaturezaDespesa, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&NaturezaDespesa{}).Where("cod_natureza_despesa = ?", codNaturezaDespesa).Updates(
		NaturezaDespesa{
			Descricao: naturezaDespesa.Descricao}).Error

	if db.Error != nil {
		return &NaturezaDespesa{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&NaturezaDespesa{}).Where("cod_natureza_despesa = ?", codNaturezaDespesa).Take(&naturezaDespesa).Error
	if err != nil {
		return &NaturezaDespesa{}, err
	}

	// retorna o elemento que foi alterado
	return naturezaDespesa, err
}

/*  =========================
	FUNCAO DELETAR NATUREZA_DESPESA POR ID
=========================  */

func (naturezaDespesa *NaturezaDespesa) DeleteNaturezaDespesa(db *gorm.DB, codNaturezaDespesa uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&NaturezaDespesa{}).Where("cod_natureza_despesa = ?", codNaturezaDespesa).Take(&NaturezaDespesa{}).Delete(&NaturezaDespesa{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("NaturezaDespesa not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
