package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR UACOM NO BANCO DE DADOS
=========================  */

func (uacom *Uacom) SaveUacom(db *gorm.DB) (*Uacom, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&uacom).Error
	if err != nil {
		return &Uacom{}, err
	}

	return uacom, nil
}

/*  =========================
	FUNCAO LISTAR UACOM POR ID
=========================  */

func (uacom *Uacom) FindUacomByID(db *gorm.DB, codIbge uint64, data string) (*Uacom, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Uacom{}).Where("cod_ibge = ? AND data = ?", codIbge, data).Take(&uacom).Error

	if err != nil {
		return &Uacom{}, err
	}

	return uacom, err
}

/*  =========================
	FUNCAO LISTAR TODAS UACOM
=========================  */

func (uacom *Uacom) FindAllUacom(db *gorm.DB) (*[]Uacom, error) {

	allUacom := []Uacom{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Uacom{}).Find(&allUacom).Error
	if err != nil {
		return &[]Uacom{}, err
	}
	return &allUacom, err
}

/*  =========================
	FUNCAO EDITAR UACOM
=========================  */

func (uacom *Uacom) UpdateUacom(db *gorm.DB, codIbge uint64, data string) (*Uacom, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE uacom SET titulo = ?, relato = ? WHERE cod_ibge = ? AND data = ?", uacom.Titulo, uacom.Relato, codIbge, data).Error

	if db.Error != nil {
		return &Uacom{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Uacom{}).Where("cod_ibge = ? AND data = ?", codIbge, data).Take(&uacom).Error
	if err != nil {
		return &Uacom{}, err
	}

	// retorna o elemento que foi alterado
	return uacom, err
}

/*  =========================
	FUNCAO DELETAR UACOM
=========================  */

func (uacom *Uacom) DeleteUacom(db *gorm.DB, codIbge uint64, data string) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Uacom{}).Where("cod_ibge = ? AND data = ?", codIbge, data).Take(&Uacom{}).Delete(&Uacom{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Uacom not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
