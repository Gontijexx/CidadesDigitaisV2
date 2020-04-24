package models

import "github.com/jinzhu/gorm"

/*  =========================
	STRUCT MODULO
=========================  */

type Modulo struct {
	CodModulo  uint32 `gorm:"primary_key;not null" json:"cod_modulo"`
	Categoria1 string `gorm:"default:null" json:"categoria_1"`
	Categoria2 string `gorm:"default:null" json:"categoria_2"`
	Categoria3 string `gorm:"default:null" json:"categoria_3"`
	Descricao  string `gorm:"default:null" json:"descricao"`
}

/*  =========================
	FUNCAO LISTAR TODOS MODULOS
=========================  */

func (modulo *Modulo) FindAllModulo(db *gorm.DB) (*[]Modulo, error) {

	allModulo := []Modulo{}

	err := db.Debug().Model(&Modulo{}).Find(&allModulo).Error
	if err != nil {
		return &[]Modulo{}, err
	}

	return &allModulo, err
}
