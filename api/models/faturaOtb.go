package models

import "github.com/jinzhu/gorm"

/*  =========================
	FUNCAO SALVAR FATURA OTB NO BANCO DE DADOS
=========================  */

func (faturaOTB *FaturaOTB) SaveFaturaOTB(db *gorm.DB) (*FaturaOTB, error) {

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

func (faturaOTB *FaturaOTB) FindFaturaOTB(db *gorm.DB, codOTB uint64) (*[]FaturaOTB, error) {

	allFaturaOTB := []FaturaOTB{}

	//	Busca todos elementos contidos no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&FaturaOTB{}).Find(&allFaturaOTB).Where("cod_otb = ?", codOTB).Error
	if err != nil {
		return &[]FaturaOTB{}, err
	}

	return &allFaturaOTB, err
}
