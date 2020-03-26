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
	FUNCAO LISTAR FATURA OTB POR ID
=========================  */

func (faturaOTB *FaturaOTB) FindFaturaOTB(db *gorm.DB, codOtb, numNF, codIbge uint64) (*FaturaOTB, error) {

	//	Busca todos elementos contidos no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(FaturaOTB{}).Where("cod_otb = ? ANF num_nf = ? AND cod_ibge = ?", codOtb, numNF, codIbge).Take(&faturaOTB).Error
	if err != nil {
		return &FaturaOTB{}, err
	}

	return faturaOTB, err
}
