package models

import "github.com/jinzhu/gorm"

/*  =========================
	FUNCAO LISTAR TODAS FATURA OTB
=========================  */

func (faturaOTB *FaturaOTB) FindFaturaOTB(db *gorm.DB, numNF, codIBGE uint64) (*[]FaturaOTB, error) {

	allFaturaOTB := []FaturaOTB{}

	//	Busca todos elementos contidos no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&FaturaOTB{}).Where("num_nf = ? AND cod_ibge = ?", numNF, codIBGE).Find(&allFaturaOTB).Error
	if err != nil {
		return &[]FaturaOTB{}, err
	}

	return &allFaturaOTB, err
}
