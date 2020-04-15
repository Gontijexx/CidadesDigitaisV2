package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	STRUCT PID
=========================  */

type Pid struct {
	CodPid  uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_pid"`
	CodIbge uint64 `gorm:"foreign_key:CodIbge;not null" json:"cod_ibge"`
	Nome    string `gorm:"not null" json:"nome"`
	Inep    string `gorm:"default:null" json:"inep"`
}

/*  =========================
	FUNCAO SALVAR PID
=========================  */

func (pid *Pid) SavePID(db *gorm.DB) (*Pid, error) {

	err := db.Debug().Create(&pid).Error
	if err != nil {
		return &Pid{}, err
	}

	return pid, err
}

/*  =========================
	FUNCAO EDITAR PID
=========================  */

func (pid *Pid) UpdatePID(db *gorm.DB, codPid uint64) (*Pid, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE pid SET cod_ibge = ?, nome = ?, inep = ? WHERE cod_pid = ?", pid.CodIbge, pid.Nome, pid.Inep, codPid)
	if db.Error != nil {
		return &Pid{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Pid{}).Where("cod_pid = ?", codPid).Take(&pid).Error
	if err != nil {
		return &Pid{}, err
	}

	// retorna o elemento que foi alterado
	return pid, err
}

/*  =========================
	FUNCAO DELETAR PID
=========================  */

func (pid *Pid) DeletePID(db *gorm.DB, codPid uint64) (int64, error) {

	db = db.Debug().Model(&Pid{}).Where("cod_pid = ?", codPid).Take(&Pid{}).Delete(&Pid{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Pid not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
