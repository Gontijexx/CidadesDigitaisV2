package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PREFEITOS NO BANCO DE DADOS
=========================  */

func (prefeitos *Prefeitos) SavePrefeitos(db *gorm.DB) (*Prefeitos, error) {

	//	Adiciona um elemento ao banco de dados
	err := db.Debug().Create(&prefeitos).Error
	if err != nil {
		return &Prefeitos{}, err
	}
	return prefeitos, nil
}

/*  =========================
	FUNCAO LISTAR PREFEITOS POR ID
=========================  */

func (prefeitos *Prefeitos) FindPrefeitosByID(db *gorm.DB, codPrefeito uint64) (*Prefeitos, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Prefeitos{}).Where("cod_prefeito = ?", codPrefeito).Take(&prefeitos).Error
	if err != nil {
		return &Prefeitos{}, err
	}
	return prefeitos, err
}

/*  =========================
	FUNCAO LISTAR TODOS PREFEITOS
=========================  */

func (prefeitos *Prefeitos) FindAllPrefeitos(db *gorm.DB) (*[]Prefeitos, error) {

	allPrefeitos := []Prefeitos{}

	//	Busca todos os elementos contidos no banco de dados
	err := db.Debug().Model(&Prefeitos{}).Find(&allPrefeitos).Error
	if err != nil {
		return &[]Prefeitos{}, err
	}

	return &allPrefeitos, err
}

/*  =========================
	FUNCAO EDITAR PREFEITOS
=========================  */

func (prefeitos *Prefeitos) UpdatePrefeitos(db *gorm.DB, codPrefeito uint64) (*Prefeitos, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE prefeitos SET cod_ibge = ?, nome = ?, cpf = ?, rg = ?, partido = ?, exercicio = ? WHERE cod_prefeito = ?", prefeitos.CodIbge, prefeitos.Nome, prefeitos.Cpf, prefeitos.RG, prefeitos.Partido, prefeitos.Exercicio, codPrefeito)
	if db.Error != nil {
		return &Prefeitos{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Prefeitos{}).Where("cod_prefeito = ?", codPrefeito).Take(&prefeitos).Error
	if err != nil {
		return &Prefeitos{}, err
	}

	//	retorna o elemento que foi alterado
	return prefeitos, err
}

/*  =========================
	FUNCAO DELETAR PREFEITOS POR ID
=========================  */

func (prefeitos *Prefeitos) DeletePrefeitos(db *gorm.DB, codPrefeito uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Prefeitos{}).Where("cod_prefeito = ?", codPrefeito).Take(&Prefeitos{}).Delete(&Prefeitos{})
	if err != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Prefeito not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
