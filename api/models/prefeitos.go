package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PREFEITOS NO BANCO DE DADOS
=========================  */

func (prefeito *Prefeito) SavePrefeito(db *gorm.DB) (*Prefeito, error) {

	//	Adiciona um elemento ao banco de dados
	err := db.Debug().Create(&prefeito).Error
	if err != nil {
		return &Prefeito{}, err
	}
	return prefeito, nil
}

/*  =========================
	FUNCAO SALVAR PREFEITOS POR ID
=========================  */

func (prefeito *Prefeito) FindPrefeitoByID(db *gorm.DB, codPrefeito uint64) (*Prefeito, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Prefeito{}).Where("cod_prefeito = ?", codPrefeito).Take(&prefeito).Error
	if err != nil {
		return &Prefeito{}, err
	}
	return prefeito, err
}

/*  =========================
	FUNCAO LISTAR TODOS PREFEITOS
=========================  */

func (prefeito *Prefeito) FindAllPrefeito(db *gorm.DB) (*[]Prefeito, error) {

	allPrefeito := []Prefeito{}

	//	Busca todos os elementos contidos no banco de dados
	err := db.Debug().Model(&Prefeito{}).Limit(100).Find(&allPrefeito).Error
	if err != nil {
		return &[]Prefeito{}, err
	}

	return &allPrefeito, err
}

/*  =========================
	FUNCAO EDITAR PREFEITO
=========================  */

func (prefeito *Prefeito) UpdatePrefeito(db *gorm.DB, codPrefeito uint64) (*Prefeito, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Prefeito{}).Where("cod_prefeito = ?", codPrefeito).Updates(
		Prefeito{
			Nome:      prefeito.Nome,
			Cpf:       prefeito.Cpf,
			RG:        prefeito.RG,
			Partido:   prefeito.Partido,
			Exercicio: prefeito.Exercicio}).Error

	if db.Error != nil {
		return &Prefeito{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Prefeito{}).Where("cod_prefeito = ?", codPrefeito).Take(&prefeito).Error
	if err != nil {
		return &Prefeito{}, err
	}

	//	retorna o elemento que foi alterado
	return prefeito, err
}

/*  =========================
	FUNCAO DELETAR PREFEITO POR ID
=========================  */

func (prefeito *Prefeito) DeletePrefeito(db *gorm.DB, codPrefeito uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Prefeito{}).Where("cod_prefeito = ?", codPrefeito).Take(&Prefeito{}).Delete(&Prefeito{})
	if err != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Prefeito not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
