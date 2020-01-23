package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR CD
=========================  */

func (cd *Cd) SaveCd(db *gorm.DB) (*Cd, error) {

	err := db.Debug().Create(&cd).Error
	if err != nil {
		return &Cd{}, err
	}
	return cd, nil

}

/*  =========================
	FUNCAO LISTAR CD POR ID
=========================  */

func (cd *Cd) FindCdByID(db *gorm.DB, cdID uint64) (*Cd, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Cd{}).Where("cod_ibge = ?", cdID).Take(&cd).Error

	if err != nil {
		return &Cd{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cd{}, errors.New("Cd Not Found")
	}

	return cd, err
}

/*  =========================
	FUNCAO LISTAR ENTIDADES
=========================  */

func (cd *Cd) FindCds(db *gorm.DB) (*[]Cd, error) {

	entity := []Cd{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Cd{}).Limit(100).Find(&entity).Error
	if err != nil {
		return &[]Cd{}, err
	}
	return &entity, err
}

/*  =========================
	FUNCAO EDITAR CD
=========================  */

func (cd *Cd) UpdateCd(db *gorm.DB, cdID uint64) (*Cd, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Cd{}).Where("cod_ibge = ?", cdID).Take(&Cd{}).UpdateColumns(
		map[string]interface{}{
			"cod_lote": cd.Cod_lote,
			"os_pe":    cd.Os_pe,
			"data_pe":  cd.Data_pe,
			"os_imp":   cd.Os_imp,
			"data_imp": cd.Data_imp,
		},
	)

	if db.Error != nil {
		return &Cd{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Cd{}).Where("cod_ibge = ?", cdID).Take(&cd).Error
	if err != nil {
		return &Cd{}, err
	}

	// retorna o elemento que foi alterado
	return cd, err
}
