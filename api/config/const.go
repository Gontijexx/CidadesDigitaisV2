package config

/*  =========================
	Server HTTP
=========================  */

const (
	// defines ip and port address for server instance
	SERVER_ADDR = "localhost:8080"
)

/*  =========================
	USER PATHS HTTP
=========================  */

const (
	USER_PATH            = "/read/usuario"
	USER_ID_PATH         = "/read/usuario/{cod_usuario}"
	USER_PATH_LOGIN      = "/read/usuario/login"
	USER_PATH_CREATEUSER = "/read/usuario/createuser"
	USER_PATH_DELETEUSER = "/read/usuario/deleteuser"
)

/*  =========================
	MODULOS PATHS HTTP
=========================  */

const (
	MODULO_PATH          = "/read/modulo"
	MODULO_USERLIST_PATH = "/read/usuario/{cod_usuario}/modulo"
)

/*  =========================
	ENTIDADE PATHS HTTP OK!
=========================  */

const (
	ENTIDADE_PATH    = "/read/entidade"
	ENTIDADE_ID_PATH = "/read/entidade/{cnpj}"
)

/*  =========================
	CONTATO PATHS HTTP OK!
=========================  */

const (
	CONTATO_PATH    = "/read/contato"
	CONTATO_ID_PATH = "/read/contato/{cod_contato}"
)

/*  =========================
	TELEFONE PATHS HTTP OK!
=========================  */

const (
	TELEFONE_PATH    = "/read/telefone"
	TELEFONE_ID_PATH = "/read/telefone/{cod_telefone}"
)

/*  =========================
	CD PATHS HTTP OK!
=========================  */

const (
	CD_PATH    = "/read/cd"
	CD_ID_PATH = "/read/cd/{cod_ibge}"
)

/*  =========================
	CD ITENS PATHS HTTP OK!
=========================  */

const (
	CD_ITENS_PATH    = "/read/cdItens"
	CD_ITENS_ID_PATH = "/read/cdItens/{cod_ibge}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	CD PROCESSO PATHS HTTP
=========================  */

const (
	PROCESSO_PATH    = "/read/processo"
	PROCESSO_ID_PATH = "/read/processo/{cod_processo}/{cod_ibge}"
)

/*  =========================
	CD UACOM PATHS HTTP
=========================  */

const (
	UACOM_PATH    = "/read/uacom"
	UACOM_ID_PATH = "/read/uacom/{cod_ibge}/{data}"
)

/*  =========================
	CD UACOM_ASSUNTO PATHS HTTP
=========================  */

const (
	UACOM_ASSUNTO_PATH = "/read/uacomAssunto"
)

/*  =========================
	CD PONTO PATHS HTTP
=========================  */

const (
	PONTO_PATH    = "/read/ponto"
	PONTO_ID_PATH = "/read/ponto/{cod_ponto}/{cod_categoria}/{cod_ibge}"
)

/*  =========================
	CD PID_TIPOOGIA PATHS HTTP
=========================  */

const (
	PID_TOPOLOGIA_ID_PATH = "/read/pontoTipologia/{cod_ponto}/{cod_categoria}/{cod_ibge}/{cod_tipologia}"
)

/*  =========================
	LOTE PATHS HTTP OK!
=========================  */

const (
	LOTE_PATH    = "/read/lote"
	LOTE_ID_PATH = "/read/lote/{cod_lote}"
)

/*  =========================
	REAJUSTE PATHS HTTP OK!
=========================  */
const (
	REAJUSTE_PATH    = "/read/reajuste"
	REAJUSTE_ID_PATH = "/read/reajuste/{cod_lote}/{ano_ref}"
)

/*  =========================
	LOTE ITENS PATHS HTTP OK!
=========================  */

const (
	LOTE_ITENS_PATH    = "/read/loteItens"
	LOTE_ITENS_ID_PATH = "/read/loteItens/{cod_lote}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	EMPENHO PATHS HTTP
=========================  */

const (
	EMPENHO_PATH    = "/read/empenho"
	EMPENHO_ID_PATH = "/read/empenho/{id_empenho}"
)

/*  =========================
	ITENS EMPENHO PATHS HTTP
=========================  */

const (
	ITENS_EMPENHO_PATH    = "/read/itensEmpenho"
	ITENS_EMPENHO_ID_PATH = "/read/itensEmpenho/{id_empenho}/{cod_item}/{cod_tipo_item}"
)

/*	=========================
	OTB PATH HTTP (PAGAMENTO)
=========================	*/

const (
	OTB_PATH    = "/read/otb"
	OTB_ID_PATH = "/read/otb/{cod_otb}"
)

/*	=========================
	OTB FATURA PATH HTTP (PAGAMENTO FATURA)
=========================	*/
const (
	OTB_FATURA_PATH    = "/read/otbFatura"
	OTB_FATURA_ID_PATH = "/read/otbFatura/{cod_otb}"
)

/*	=========================
	OTB ITENS PATH HTTP (PAGAMENTO ITENS)
=========================	*/

const (
	OTB_ITENS_PATH    = "/read/otbItens"
	OTB_ITENS_ID_PATH = "/read/otbItens/{cod_otb}"
)

/*+++++++++++++++++++++++++++++++++++++++++++++++++
	FATURA_ITENS_PATH    = "/read/faturaItens"
	FATURA_ITENS_LIST    = "/read/faturaItens/:num_nf"
	FATURA_ITENS_LIST_ID = "/read/faturaItens/:municipio_cod_ibge/:natureza_despesa_cod_natureza_despesa/:cod_item/:cod_tipo_item"
	FATURA_ITENS_DELETE  = "/read/faturaItens/:fatura_num_nf/:cod_empenho/:cod_item/:cod_tipo_item"
+++++++++++++++++++++++++++++++++++++++++++++++++*/

/*	=========================
	FATURA PATH HTTP
=========================	*/

const (
	FATURA_PATH    = "/read/fatura"
	FATURA_ID_PATH = "/read/fatura/{num_nf}"
)

/*	=========================
	FATURA OTB PATH HTTP (FATURA PAGAMENTO)
=========================	*/
const (
	FATURA_OTB_ID_PATH = "/read/faturaOtb/{num_nf}/{cod_ibge}"
)

/*	=========================
	FATURA ITENS PATH HTTP (FATURA ITENS)
=========================	*/

const (
	FATURA_ITENS_PATH    = "/read/faturaItens"
	FATURA_ITENS_ID_PATH = "/read/faturaItens/{cod_otb}"
)

/*  =========================
	PREVISAO EMPENHO PATHS HTTP OK!
=========================  */

const (
	PREVISAO_EMPENHO_PATH    = "/read/previsaoEmpenho"
	PREVISAO_EMPENHO_ID_PATH = "/read/previsaoEmpenho/{cod_previsao_empenho}"
)

/*  =========================
	ITENS PREVISAO EMPENHO PATHS HTTP
=========================  */

const (
	ITENS_PREVISAO_EMPENHO_PATH    = "/read/itensPrevisaoEmpenho"
	ITENS_PREVISAO_EMPENHO_ID_PATH = "/read/itensPrevisaoEmpenho/{cod_previsao_empenho}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	ASSUNTO PATHS HTTP OK!
=========================  */

const (
	ASSUNTO_PATH    = "/read/assunto"
	ASSUNTO_ID_PATH = "/read/assunto/{cod_assunto}"
)

/*	=========================
	CATEGORIA PATHS HTTP OK!
=========================	*/
const (
	CATEGORIA_PATH    = "/read/categoria"
	CATEGORIA_ID_PATH = "/read/categoria/{cod_categoria}"
)

/*	=========================
	CLASSE EMPENHO PATHS HTTP OK!
=========================	*/

const (
	CLASSE_EMPENHO_PATH    = "/read/classeEmpenho"
	CLASSE_EMPENHO_ID_PATH = "/read/classeEmpenho/{cod_classe_empenho}"
)

/*  =========================
	ETAPA PATHS HTTP
=========================  */

const (
	ETAPA_PATH    = "/read/etapa"
	ETAPA_ID_PATH = "/read/etapa/{cod_etapa}"
)

/*	=========================
	ITENS PATHS HTTP OK!
=========================	*/

const (
	ITENS_PATH    = "/read/itens"
	ITENS_ID_PATH = "/read/itens/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	MUNICIPIO PATHS HTTP OK!
=========================  */

const (
	MUNICIPIO_PATH    = "/read/municipio"
	MUNICIPIO_ID_PATH = "/read/municipio/{cod_ibge}"
)

/*	=========================
	NATUREZA DE DESPESA PATHS HTTP OK!
=========================	*/

const (
	NATUREZA_DESPESA_PATH    = "/read/naturezaDespesa"
	NATUREZA_DESPESA_ID_PATH = "/read/naturezaDespesa/{cod_natureza_despesa}"
)

/*  =========================
	PREFEITOS PATHS HTTP OK!
=========================  */

const (
	PREFEITOS_PATH    = "/read/prefeitos"
	PREFEITOS_ID_PATH = "/read/prefeitos/{cod_prefeito}"
)

/*	=========================
	TIPOLOGIAS PATHS HTTP OK!
=========================	*/

const (
	TIPOLOGIA_PATH    = "/read/tipologia"
	TIPOLOGIA_ID_PATH = "/read/tipologia/{cod_tipologia}"
)

/*	=========================
	TIPO ITEM PATHS HTTP OK!
=========================	*/

const (
	TIPO_ITEM_PATH    = "/read/tipoItem"
	TIPO_ITEM_ID_PATH = "/read/tipoItem/{cod_tipo_item}"
)
