package models

/*  =========================
	TABELA TELEFONE
=========================  */

type Telefone struct {
	CodTelefone uint64 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_telefone"`
	CodContato  uint64 `gorm:"foreign_key:CodContato;not null;size:11" json:"cod_contato"`
	Telefone    string `gorm:"default:null;size:11" json:"telefone"`
	Tipo        string `gorm:"default:null;size:10" json:"tipo"`
}

/*  =========================
	TABELA CD PROCESSO
=========================  */

type Processo struct {
	CodProcesso string `gorm:"primary_key;not null;size:17" json:"cod_processo"`
	CodIbge     uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Descricao   string `gorm:"default:null" json:"descricao"`
}

/*  =========================
	TABELA UACOM (CD)
=========================  */

type Uacom struct {
	CodIbge uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Data    string `gorm:"primary_key;not null" json:"data"`
	Titulo  string `gorm:"default:null" json:"titulo"`
	Relato  string `gorm:"default:null" json:"relato"`
}

/*  =========================
	TABELA UACOM ASSUNTO (CD)
=========================  */

type UacomAssunto struct {
	CodIbge    uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Data       string `gorm:"primary_key;foreign_key:Data;not null" json:"data"`
	CodAssunto uint64 `gorm:"primary_key;foreign_key:CodAssunto;not null" json:"cod_assunto"`
}

/*  =========================
	TABELA PONTO (CD)
=========================  */

type Ponto struct {
	CodPonto     uint64 `gorm:"primary_key;not null" json:"cod_ponto"`
	CodCategoria uint64 `gorm:"primary_key;foreign_key:CodCategoria;not null" json:"cod_categoria"`
	CodIbge      uint64 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	CodPid       uint64 `gorm:"foreign_key:CodPid;not null" json:"cod_pid"`
	Endereco     string `gorm:"default:null" json:"endereco"`
	Numero       string `gorm:"default:null;size:10" json:"numero"`
	Complemento  string `gorm:"default:null" json:"complemento"`
	Bairro       string `gorm:"default:null" json:"bairro"`
	Cep          string `gorm:"default:null;size:8" json:"cep"`
	Latitude     uint64 `gorm:"default:null" json:"latitude"`
	Longitude    uint64 `gorm:"default:null" json:"longitude"`
}

/*  =========================
	TABELA REAJUSTE
=========================  */

type Reajuste struct {
	AnoRef     uint64  `gorm:"primary_key;not null;size:11" json:"ano_ref"`
	CodLote    uint64  `gorm:"primary key;foreign_key:CodLote;not null;size:11" json:"cod_lote"`
	Percentual float64 `gorm:"default:null" json:"percentual"`
}

/*  =========================
	TABELA PREVISAO EMPENHO
=========================  */

type PrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint64 `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;auto_increment;not null" json:"cod_previsao_empenho"`
	CodLote            uint64 `gorm:"foreign_key:CodLote;not null;size:11" json:"cod_lote"`
	CodNaturezaDespesa uint64 `gorm:"foreign_key:CodNaturezaDespesa;not null;size:11" json:"cod_natureza_despesa"`
	Data               string `gorm:"default:null" json:"data"`
	Tipo               string `gorm:"default:null;size:1" json:"tipo"`
	Ano_referencia     uint64 `gorm:"default:null;size:11" json:"ano_referencia"`
}

/*  =========================
	TABELA PREFEITOS
=========================  */

type Prefeitos struct {
	CodPrefeito uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_prefeito"`
	CodIbge     uint64 `gorm:"foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Nome        string `gorm:"default:null" json:"nome"`
	Cpf         string `gorm:"default:null" json:"cpf"`
	RG          string `gorm:"default:null" json:"rg"`
	Partido     string `gorm:"default:null" json:"partido"`
	Exercicio   string `gorm:"default:null" json:"exercicio"`
}

/*	=========================
		TABELA TIPOLOGIAS
=========================	*/

type Tipologia struct {
	CodTipologia uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_tipologia"`
	Descricao    string `gorm:"default:null" json:"descricao"`
}

/*	=========================
		TABELA TIPO ITEM
=========================	*/

type TipoItem struct {
	CodTipoItem uint64 `gorm:"primary_key;not null" json:"cod_tipo_item"`
	Descricao   string `gorm:"default:null" json:"descricao"`
}
