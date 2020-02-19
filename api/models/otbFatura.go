package models

import "github.com/jinzhu/gorm"

/*  =========================
	FUNCAO SALVAR FATURA OTB NO BANCO DE DADOS
=========================  */

func (faturaOTB *FaturaOTB) SaveOTBFatura(db *gorm.DB) (*FaturaOTB, error) {

	//	Adiciona um novo elemento no banco de dados
	err := db.Debug().Create(&faturaOTB).Error
	if err != nil {
		return &FaturaOTB{}, err
	}

	return faturaOTB, nil

}

/*  =========================
	FUNCAO LISTAR TODAS FATURA OTB
=========================  */

func (faturaOTB *FaturaOTB) FindOTBFatura(db *gorm.DB, codOTB uint64) (*[]FaturaOTB, error) {

	allFaturaOTB := []FaturaOTB{}

	//	Busca todos elementos contidos no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&FaturaOTB{}).Where("cod_otb = ?", codOTB).Find(&allFaturaOTB).Error
	if err != nil {
		return &[]FaturaOTB{}, err
	}

	return &allFaturaOTB, err
}
