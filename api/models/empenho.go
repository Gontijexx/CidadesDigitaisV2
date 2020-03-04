package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR EMPENHO NO BANCO DE DADOS
=========================  */

func (empenho *Empenho) SaveEmpenho(db *gorm.DB) (*Empenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&empenho).Error
	if err != nil {
		return &Empenho{}, err
	}
	return empenho, nil
}

/*  =========================
	FUNCAO LISTAR EMPENHO POR ID
=========================  */

func (empenho *Empenho) FindEmpenhoByID(db *gorm.DB, idEmpenho uint64) (*Empenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Empenho{}).Where("id_empenho = ?", idEmpenho).Take(&empenho).Error

	if err != nil {
		return &Empenho{}, err
	}

	return empenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS EMPENHO
=========================  */

func (empenho *Empenho) FindAllEmpenho(db *gorm.DB) (*[]Empenho, error) {

	allEmpenho := []Empenho{}

	//	Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Empenho{}).Find(&allEmpenho).Error
	if err != nil {
		return &[]Empenho{}, err
	}
	return &allEmpenho, err
}

/*  =========================
	FUNCAO EDITAR EMPENHO
=========================  */

func (empenho *Empenho) UpdateEmpenho(db *gorm.DB, idEmpenho uint64) (*Empenho, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Empenho{}).Where("id_empenho = ?", idEmpenho).Updates(
		Empenho{
			CodPrevisaoEmpenho: empenho.CodPrevisaoEmpenho,
			CodEmpenho:         empenho.CodEmpenho,
			Data:               empenho.Data,
			Contador:           empenho.Contador}).Error

	if err != nil {
		return &Empenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Empenho{}).Where("id_empenho = ?", idEmpenho).Take(&empenho).Error
	if err != nil {
		return &Empenho{}, err
	}

	//	Retorna o elemento que foi alterado
	return empenho, err
}

/*  =========================
	FUNCAO DELETAR EMPENHO
=========================

func (empenho *Empenho) DeleteEmpenho(db *gorm.DB, idEmpenho uint64) (int64, error) {

	//	Deleta  um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Empenho{}).Where("id_empenho = ?", idEmpenho).Take(&Empenho{}).Delete(&Empenho{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Empenho not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

*/
