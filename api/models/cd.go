package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/* =========================
	STRUCT CD
=========================  */

type CD struct {
	CodIbge       uint32 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	CodLote       uint32 `gorm:"foreign_key:CodLote;not null" json:"cod_lote"`
	NomeMunicipio string `gorm:"default:null" json:"nome_municipio"`
	OsPe          string `gorm:"size:10;default:null" json:"os_pe"`
	DataPe        string `gorm:"default:null" json:"data_pe"`
	OsImp         string `gorm:"size:10;default:null" json:"os_imp"`
	DataImp       string `gorm:"default:null" json:"data_imp"`
}

/*  =========================
	FUNCAO SALVAR CD
=========================  */

func (cd *CD) SaveCD(db *gorm.DB) (*CD, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&cd).Error
	if err != nil {
		return &CD{}, err
	}

	return cd, err
}

/*  =========================
	FUNCAO LISTAR CD POR ID
=========================  */

func (cd *CD) FindCDByID(db *gorm.DB, codIbge uint32) (*CD, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(CD{}).Where("cod_ibge = ?", codIbge).Take(&cd).Error
	if err != nil {
		return &CD{}, err
	}

	return cd, err
}

/*  =========================
	FUNCAO LISTAR CD
=========================  */

func (cd *CD) FindAllCD(db *gorm.DB) (*[]CD, error) {

	allCD := []CD{}

	// Busca todos elementos contidos no banco de dados e faz join com a tabela municipio
	err := db.Debug().Table("cd").Select("CONCAT(municipio.nome_municipio, ' - ', municipio.uf) AS nome_municipio, cd.*").
		Joins("JOIN municipio ON cd.cod_ibge = municipio.cod_ibge").Scan(&allCD).Error
	if err != nil {
		return &[]CD{}, err
	}

	return &allCD, err
}

/*  =========================
	FUNCAO EDITAR CD
=========================  */

func (cd *CD) UpdateCD(db *gorm.DB, codIbge uint32) (*CD, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE cd SET cod_lote = ?, os_pe = ?, data_pe = ?, os_imp = ?, data_imp = ? WHERE cod_ibge = ?", cd.CodLote, cd.OsPe, cd.DataPe, cd.OsImp, cd.DataImp, codIbge)
	if db.Error != nil {
		return &CD{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&CD{}).Where("cod_ibge = ?", codIbge).Take(&cd).Error
	if err != nil {
		return &CD{}, err
	}

	return cd, err
}

/*  =========================
	FUNCAO DELETAR CD POR ID
=========================	*/

func (cd *CD) DeleteCD(db *gorm.DB, codIbge uint32) error {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&CD{}).Where("cod_ibge = ?", codIbge).Take(&CD{}).Delete(&CD{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("CD not found")
		}
		return db.Error
	}

	return db.Error
}
