package models

/*  =========================
	TABELA ENTIDADE
=========================  */

type Entidade struct {
	Cnpj          string `gorm:"primary_key;not null;size:14" json:"cnpj" validate:"alphanum":`
	Nome          string `gorm:"size:50;default:null" json:"nome" validate:"alphanum":`
	Endereco      string `gorm:"size:100;default:null" json:"endereco" validate:"alphanum":`
	Numero        string `gorm:"size:10;default:null" json:"numero" validate:"alphanum":`
	Bairro        string `gorm:"size:100;default:null" json:"bairro" validate:"alphanum":`
	Cep           string `gorm:"size:8;default:null" json:"cep" validate:"alphanum":`
	NomeMunicipio string `gorm:"size:50;default:null" json:"nome_municipio" validate:"alphanum":`
	UF            string `gorm:"size:2;default:null" json:"uf" validate:"alphanum":`
	Observacao    string `gorm:"size:1000;default:null" json:"observacao" validate:"alphanum":`
}

/*  =========================
	TABELA CONTATO
=========================  */

type Contato struct {
	CodContato uint64 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_contato" validate:"required":`
	Cnpj       string `gorm:"foreign_key:Cnpj;not null;size:14" json:"cnpj" validate:"required":`
	CodIbge    uint64 `gorm:"foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"alphanum":`
	Nome       string `gorm:"default:null;size:50" json:"nome" validate:"alphanum":`
	Email      string `gorm:"default:null;size:100" json:"email" validate:"email":`
	Funcao     string `gorm:"default:null;size:45" json:"funcao" validate:"alphanum":`
}

/*  =========================
	TABELA TELEFONE
=========================  */

type Telefone struct {
	CodTelefone uint64 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_telefone" validate:"number":`
	CodContato  uint64 `gorm:"foreign_key:CodContato;not null;size:11" json:"cod_contato" validate:"number":`
	Telefone    string `gorm:"default:null;size:11" json:"telefone" validate:"required":`
	Tipo        string `gorm:"default:null;size:10" json:"tipo" validate:"alphanum":`
}

/*  =========================
	TABELA CD
=========================  */

type CD struct {
	CodIbge uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"number":`
	CodLote uint64 `gorm:"foreign_key:CodLote;not null;size:11" json:"cod_lote" validate:"number":`
	OsPe    string `gorm:"size:10;default:null" json:"os_pe" validate:"alphanum":`
	DataPe  string `gorm:"default:null" json:"data_pe" validate:"alphanum":`
	OsImp   string `gorm:"size:10;default:null" json:"os_imp" validate:"alphanum":`
	DataImp string `gorm:"default:null" json:"data_imp" validate:"alphanum":`
}

/*  =========================
	TABELA CD ITENS
=========================  */

type CDItens struct {
	CodIbge                    uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"number":`
	CodItem                    uint64 `gorm:"primary_key;foreign_key:CodItem;not null;size:11" json:"cod_item" validate:"number":`
	CodTipoItem                uint64 `gorm:"primary_key;foreign_key:CodTipo_item;not null;size:11" json:"cod_tipo_item" validate:"number":`
	QuantidadePrevisto         uint64 `gorm:"default:null;size:11" json:"quantidade_previsto" validate:"required":`
	QuantidadeProjetoExecutivo uint64 `gorm:"default:null;size:11" json:"quantidade_projeto_executivo" validate:"alphanum":`
	QuantidadeTermoInstalacao  uint64 `gorm:"default:null;size:11" json:"quantidade_termo_instalacao" validate:"alphanum":`
}

/*  =========================
	TABELA CD PROCESSO
=========================  */

type Processo struct {
	CodProcesso string `gorm:"primary_key;not null;size:17" json:"cod_processo" validate:"alphanum"`
	CodIbge     uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"number"`
	Descricao   string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}

/*  =========================
	TABELA UACOM (CD)
=========================  */

type Uacom struct {
	CodIbge uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"number"`
	Data    string `gorm:"primary_key;not null" json:"data" validate:"alphanum"`
	Titulo  string `gorm:"default:null" json:"titulo" validate:"alphanum"`
	Relato  string `gorm:"default:null" json:"relato" validate:"alphanum"`
}

/*  =========================
	TABELA UACOM ASSUNTO (CD)
=========================  */

type UacomAssunto struct {
	CodIbge    uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"number"`
	Data       string `gorm:"primary_key;foreign_key:Data;not null" json:"data" validate:"alphanum"`
	CodAssunto uint64 `gorm:"primary_key;foreign_key:CodAssunto;not null" json:"cod_assunto" validate:"number"`
}

/*  =========================
	TABELA PONTO (CD)
=========================  */

type Ponto struct {
	CodPonto     uint64 `gorm:"primary_key;not null" json:"cod_ponto" validate:"number"`
	CodCategoria uint64 `gorm:"primary_key;foreign_key:CodCategoria;not null" json:"cod_categoria" validate:"number"`
	CodIbge      uint64 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge" validate:"number"`
	CodPID       uint64 `gorm:"foreign_key:CodPID;not null" json:"cod_pid" validate:"number"`
	Endereco     string `gorm:"default:null" json:"endereco" validate:"alphanum"`
	Numero       string `gorm:"default:null;size:10" json:"numero" validate:"alphanum"`
	Complemento  string `gorm:"default:null" json:"complemento" validate:"alphanum"`
	Bairro       string `gorm:"default:null" json:"bairro" validate:"alphanum"`
	Cep          string `gorm:"default:null;size:8" json:"cep" validate:"alphanum"`
	Latitude     uint64 `gorm:"default:null" json:"latitude" validate:"number"`
	Longitude    uint64 `gorm:"default:null" json:"longitude" validate:"number"`
}

/*  =========================
	TABELA PID (CD)
=========================  */

type PID struct {
	CodPID  uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_pid" validate:"number"`
	CodIbge uint64 `gorm:"foreign_key:CodIbge;not null" json:"cod_ibge" validate:"number"`
	Nome    string `gorm:"default:null" json:"nome" validate:"alphanum"`
	Inep    string `gorm:"default:null" json:"inep" validate:"alphanum"`
}

/*  =========================
	TABELA PID TIPOLOGIA (CD)
=========================  */

type PIDTipologia struct {
	CodPID       uint64 `gorm:"primary_key;foreign_key:CodPID;not null" json:"cod_pid" validate:"number"`
	CodTipologia uint64 `gorm:"primary_key;foreign_key:CodTipologia;not null" json:"cod_tipologia" validate:"number"`
}

/*  =========================
	TABELA LOTE
=========================  */

type Lote struct {
	CodLote     uint64 `gorm:"primary_key;not null;size:11" json:"cod_lote" validate:"number":`
	Cnpj        string `gorm:"foreign_key:Cnpj;not null;size:14" json:"cnpj" validate:"required, number":`
	Contrato    string `gorm:"default:null;size:10" json:"contrato" validate:"alphanum":`
	DtInicioVig string `gorm:"default:null;size:10" json:"dt_inicio_vig" validate:"alphanum":`
	DtFinalVig  string `gorm:"default:null;size:10" json:"dt_final_vig" validate:"alphanum":`
	DtReajuste  string `gorm:"default:null;size:10" json:"dt_reajuste" validate:"alphanum":`
}

/*  =========================
	TABELA REAJUSTE
=========================  */

type Reajuste struct {
	AnoRef     uint64  `gorm:"primary_key;not null;size:11" json:"ano_ref" validate:"number":`
	CodLote    uint64  `gorm:"primary key;foreign_key:CodLote;not null;size:11" json:"cod_lote" validate:"number":`
	Percentual float64 `gorm:"default:null" json:"percentual" validate:"alphanum":`
}

/*  =========================
	TABELA LOTE ITENS
=========================  */

type LoteItens struct {
	CodLote     uint64  `gorm:"primary_key;foreign_key:CodLote;not null;size:11" json:"cod_lote" validate:"number":`
	CodItem     uint64  `gorm:"primary_key;foreign_key:CodItem;not null;size:11" json:"cod_item" validate:"number":`
	CodTipoItem uint64  `gorm:"primary_key;foreign_key:CodTipoItem;not null;size:11" json:"cod_tipo_item" validate:"number":`
	Preco       float64 `gorm:"default:null;size:12" json:"preco" validate:"alphanum":`
}

/*  =========================
	TABELA EMPENHO
=========================  */

type Empenho struct {
	IDEmpenho          uint64 `gorm:"primary_key;not null" json:""id_empenho validate:"number"`
	CodPrevisaoEmpenho uint64 `gorm:"foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho" validate:"number"`
	CodEmpenho         string `gorm:"default:null;size:13" json:"cod_empenho" validate:"alphanum"`
	Data               string `gorm:"default:null" json:"data" validate:"required"`
	Contador           uint64 `gorm:"default:null" json:"contador" validate:"alphanum"`
}

/*  =========================
	TABELA ITENS EMPENHO
=========================  */

type ItensEmpenho struct {
	IDEmpenho          uint64  `gorm:"primary_key;foreign_key:IDEmpenho;not null;size:13" json:"id_empenho" validate:"number"`
	CodItem            uint64  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item" validate:"number"`
	CodTipoItem        uint64  `gorm:"primary_key;foreign_key:CodTipoItem;not null" json:"cod_tipo_item" validate:"number"`
	CodPrevisaoEmpenho uint64  `gorm:"foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho" validate:"number"`
	Valor              float64 `gorm:"default:null" json:"valor" validate:"alphanum"`
	Quantidade         uint64  `gorm:"default:null" json:"quantidade" validate:"number"`
}

/*	=========================
		TABELA PAGAMENTO (OTB)
=========================	*/

type Otb struct {
	CodOtb uint64 `gorm:"primary_key;not null" json:"cod_otb" validate:"number"`
	DtPgto uint64 `gorm:"default:null" json:"dt_pgto" validate:"alphanum"`
}

/*	=========================
		TABELA FATURA (OTB)
=========================	*/

/*	=========================
		TABELA ITENS (OTB)
=========================	*/

/*	=========================
		TABELA FATURA
=========================	*/

/*	=========================
		TABELA ITENS FATURA
=========================	*/

/*  =========================
	TABELA PREVISAO EMPENHO
=========================  */

type PrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint64 `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;auto_increment;not null" json:"cod_previsao_empenho" validate:"number":`
	CodLote            uint64 `gorm:"foreign_key:CodLote;not null;size:11" json:"cod_lote" validate:"number":`
	CodNaturezaDespesa uint64 `gorm:"foreign_key:CodNaturezaDespesa;not null;size:11" json:"cod_natureza_despesa" validate:"number":`
	Data               string `gorm:"default:null" json:"data" validate:"required":`
	Tipo               string `gorm:"default:null;size:1" json:"tipo" validate:"alphanum":`
	Ano_referencia     uint64 `gorm:"default:null;size:11" json:"ano_referencia" validate:"alphanum":`
}

/*  =========================
	TABELA ITENS PREVISAO EMPENHO
=========================  */

type ItensPrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint64  `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho" validate:"number":`
	CodItem            uint64  `gorm:"primary_key;foreign_key:CodItem;not null;size:11" json:"cod_item" validate:"number":`
	CodTipoItem        uint64  `gorm:"primary_key;foreign_key:CodTipo_item;not null;size:11" json:"cod_tipo_item" validate:"number":`
	CodLote            uint64  `gorm:"foreign_key:CodLote;not null;size:11" json:"cod_lote" validate:"number":`
	Valor              float64 `gorm:"default:null;size:12" json:"valor" validate:"alphanum":`
	Quantidade         uint64  `gorm:"default:null;size:11" json:"quantidade" validate:"alphanum":`
}

/*	=========================
		TABELA ASSUNTO
=========================	*/

type Assunto struct {
	CodAssunto uint64 `gorm:"primary_key;auto_increment;not null;" json:"cod_assunto" validate:"number"`
	Descricao  string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}

/*	=========================
		TABELA CATEGORIA
=========================	*/

type Categoria struct {
	CodCategoria uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_categoria" validate:"number"`
	Descricao    string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}

/*	=========================
		TABELA CLASSE EMPENHO
=========================	*/

type ClasseEmpenho struct {
	CodClasseEmpenho uint64 `gorm:"primary_key;not null" json:"cod_classe_empenho" validate:"number"`
	Descricao        string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}

/*	=========================
		TABELA ETAPA
=========================	*/

type Etapa struct {
	CodEtapa  uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_etapa" validate:"number"`
	Descricao string `gorm:"default:null" json:"descricao" validate:"alphanum"`
	Duracao   uint64 `gorm:"default:null" json:"duracao" validate:"number"`
	Depende   uint64 `gorm:"default:null" json:"depende" validate:"number"`
	Delay     uint64 `gorm:"default:null" json:"delay" validate:"number"`
	SetorResp string `gorm:"default:null" json:"setor_resp" validate:"alphanum"`
}

/*	=========================
		TABELA ITENS
=========================	*/

type Itens struct {
	CodItem            uint64 `gorm:"primary_key;not null" json:"cod_item" validate:"number"`
	CodTipoItem        uint64 `gorm:"primary_key;foreign_key:CodTipoItem;not null" json:"cod_tipo_item" validate:"number"`
	CodNaturezaDespesa uint64 `gorm:"foreign_key:CodNaturezaDespesa" json:"cod_natureza_despesa" validate:"number"`
	CodClasseEmpenho   uint64 `gorm:"foreign_key:CodClasseEmpenho" json:"cod_classe_empenho" validate:"number"`
	Descricao          string `gorm:"default:null" json:"descricao" validate:"alphanum"`
	Unidade            string `gorm:"default:null" json:"unidade" validate:"alphanum"`
}

/*  =========================
	TABELA MUNICIPIOS
=========================  */

type Municipios struct {
	CodIbge       uint64  `gorm:"primary_key;not null;size:7" json:"cod_ibge" validate:"number"`
	NomeMunicipio string  `gorm:"default:null" json:"nome_municipio" validate:"alphanum"`
	Populacao     uint64  `gorm:"default:null" json:"populacao" validate:"number"`
	UF            string  `gorm:"default:null;size:2" json:"uf" validate:"alphanum"`
	Regiao        string  `gorm:"default:null;size:15" json:"regiao" validate:"alphanum"`
	Cnpj          string  `gorm:"default:null;size:14" json:"cnpj" validate:"alphanum"`
	DistCapital   uint64  `gorm:"default:null" json:"dist_capital" validate:"number"`
	Endereco      string  `gorm:"default:null" json:"endereco" validate:"alphanum"`
	Numero        string  `gorm:"default:null;size:10" json:"numero" validate:"alphanum"`
	Complemento   string  `gorm:"default:null" json:"complemento" validate:"alphanum"`
	Bairro        string  `gorm:"default:null" json:"bairro" validate:"alphanum"`
	Idhm          float64 `gorm:"default:null" json:"idhm" validate:"number"`
	Latitude      float64 `gorm:"default:null" json:"latitude" validate:"number"`
	Longitude     float64 `gorm:"default:null" json:"longitude" validate:"number"`
}

/*	=========================
	TABELA NATUREZA DE DESPESA
=========================	*/

type NaturezaDespesa struct {
	CodNaturezaDespesa uint64 `gorm:"primary_key;not null" json:"cod_natureza_despesa" validate:"number"`
	Descricao          string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}

/*  =========================
	TABELA PREFEITOS
=========================  */

type Prefeito struct {
	CodPrefeito uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_prefeito" validate:"number"`
	CodIbge     uint64 `gorm:"foreign_key:CodIbge;not null;size:7" json:"cod_ibge" validate:"number"`
	Nome        string `gorm:"default:null" json:"nome" validate:"alphanum"`
	Cpf         string `gorm:"default:null" json:"cpf" validate:"alphanum"`
	RG          string `gorm:"default:null" json:"rg" validate:"alphanum"`
	Partido     string `gorm:"default:null" json:"partido" validate:"alphanum"`
	Exercicio   string `gorm:"default:null" json:"exercicio" validate:"alphanum"`
}

/*	=========================
		TABELA TIPOLOGIAS
=========================	*/

type Tipologia struct {
	CodTipologia uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_tipologia" validate:"number"`
	Descricao    string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}

/*	=========================
		TABELA TIPO ITEM
=========================	*/

type TipoItem struct {
	CodTipoItem uint64 `gorm:"primary_key;not null" json:"cod_tipo_item" validate:"number"`
	Descricao   string `gorm:"default:null" json:"descricao" validate:"alphanum"`
}
