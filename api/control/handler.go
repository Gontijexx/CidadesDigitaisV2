package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) CreateHandler() (r *mux.Router) {

	//	CRIA UM ROTEADOR

	r = s.Router

	//	HOME
	r.HandleFunc("/{id}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.Home))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM USUARIO
	=========================	*/

	//ROTA DE LOGIN
	r.HandleFunc("/read/usuario/login", middlewares.SetMiddleJSON(s.Login)).Methods(http.MethodPost)

	//LISTA USUARIOS
	r.HandleFunc("/read/usuario", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUsers))).Methods(http.MethodGet)

	//SALVA USUARIO
	r.HandleFunc("/read/usuario/createuser", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateUser))).Methods(http.MethodPost)

	//LISTA USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUser))).Methods(http.MethodGet)

	//EDITA O USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateUser))).Methods(http.MethodPut)

	//APAGA O USUARIO
	//	r.HandleFunc(config.USER_ID_PATH, middlewares.SetMiddleAuth(s.DeleteUser)).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM MODULOS
	=========================	*/

	/*	=========================
			ROTAS EM ENTIDADE OK!
	=========================	*/

	//	LISTA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllEntidade))).Methods(http.MethodGet)

	//	SALVA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateEntidade))).Methods(http.MethodPost)

	//	EDITA ENTIDADE
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEntidade))).Methods(http.MethodPut)

	//	LISTA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEntidadeByID))).Methods(http.MethodGet)

	//	APAGA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEntidade))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS DE CONTATO OK!
	=========================	*/

	//	LISTA CONTATO
	r.HandleFunc(config.CONTATO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllContato))).Methods(http.MethodGet)

	//	SALVA CONTATO
	r.HandleFunc(config.CONTATO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateContato))).Methods(http.MethodPost)

	//	EDITA CONTATO (cod_contato)
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateContato))).Methods(http.MethodPut)

	//	APAGA LOTE (cod_contato)
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteContato))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS DE TELEFONE OK!
	=========================	*/

	//	LISTA TELEFONE
	r.HandleFunc(config.TELEFONE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllTelefone))).Methods(http.MethodGet)

	//	SALVA TELEFONE
	r.HandleFunc(config.TELEFONE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateTelefone))).Methods(http.MethodPost)

	//	APAGA TELEFONE (cod_telefone)
	r.HandleFunc(config.TELEFONE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteTelefone))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM CD 13000
	=========================	*/

	//	LISTA CD
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCD))).Methods(http.MethodGet)

	//	SALVA CD
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateCD))).Methods(http.MethodPost)

	//	EDITA CD
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCD))).Methods(http.MethodPut)

	//	LISTA CD POR ID
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM CD_ITENS 13000
	=========================	*/

	//	LISTA CD_ITENS
	r.HandleFunc(config.CD_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCDItens))).Methods(http.MethodGet)

	//	EDITA CD_ITENS
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCDItens))).Methods(http.MethodPut)

	//	LISTA CD_ITENS POR ID
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDItensByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM LOTE OK!
	=========================	*/

	//	LISTA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllLote))).Methods(http.MethodGet)

	//	SALVA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateLote))).Methods(http.MethodPost)

	//	EDITA LOTE
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLote))).Methods(http.MethodPut)

	//	LISTA LOTE (cod_lote)
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteByID))).Methods(http.MethodGet)

	//	APAGA LOTE (cod_lote)
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteLote))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM REAJUSTE OK!
	=========================	*/

	//	SALVA REAJUSTE
	r.HandleFunc(config.REAJUSTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateReajuste))).Methods(http.MethodPost)

	//	LISTA REAJUSTE
	r.HandleFunc(config.REAJUSTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllReajuste))).Methods(http.MethodGet)

	//	APAGA REAJUSTE (lote_cod_lote, ano_ref)
	r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteReajuste))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM LOTE ITENS OK!
	=========================	*/

	//	LISTA LOTE
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllLoteItens))).Methods(http.MethodGet)

	//	EDITA LOTE ITENS
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLoteItens))).Methods(http.MethodPut)

	//	LISTA LOTE ITENS POR ID
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteItensByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM EMPENHO OK!
	=========================	*/

	//	LISTA EMPENHO
	r.HandleFunc(config.EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllEmpenho))).Methods(http.MethodGet)

	//	SALVA EMPENHO
	r.HandleFunc(config.EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateEmpenho))).Methods(http.MethodPost)

	//	EDITA EMPENHO (cod_empenho)
	r.HandleFunc(config.EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEmpenho))).Methods(http.MethodPut)

	//	LISTA EMPENHO (cod_empenho)
	r.HandleFunc(config.EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEmpenhoByID))).Methods(http.MethodGet)

	//	APAGA EMPENHO (cod_empenho)
	r.HandleFunc(config.EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEmpenho))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ITENS EMPENHO OK!
	=========================	*/

	//	EDITA ITENS EMPENHO (cod_empenho)
	r.HandleFunc(config.ITENS_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS EMPENHO (cod_empenho)
	r.HandleFunc(config.ITENS_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensEmpenhoByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM PAGAMENTO 16000
	=========================	*/

	/*	=========================
			ROTAS EM FATURA 17000
	=========================	*/

	/*	=========================
			ROTAS EM PREVISAO EMPENHO OK!
	=========================	*/

	//	LISTA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPrevisaoEmpenho))).Methods(http.MethodGet)

	//	SALVA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrevisaoEmpenho))).Methods(http.MethodPost)

	//	EDITA PREVISAO EMPENHO (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA PREVISAO EMPENHO (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	//	APAGA PREVISAO EMPENHO (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePrevisaoEmpenho))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ITENS PREVISAO EMPENHO OK!
	=========================	*/

	//	EDITA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensPrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM ASSUNTO OK!
	=========================	*/

	//	LISTA ASSUNTO
	r.HandleFunc(config.ASSUNTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllAssunto))).Methods(http.MethodGet)

	//	SALVA ASSUNTO
	r.HandleFunc(config.ASSUNTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateAssunto))).Methods(http.MethodPost)

	//	EDITA ASSUNTO (cod_assunto)
	r.HandleFunc(config.ASSUNTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateAssunto))).Methods(http.MethodPut)

	//	LISTA ASSUNTO (cod_assunto)
	r.HandleFunc(config.ASSUNTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAssuntoByID))).Methods(http.MethodGet)

	//	APAGA ASSUNTO (cod_assunto)
	r.HandleFunc(config.ASSUNTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteAssunto))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM CATEGORIA OK!
	=========================	*/
	//	LISTA CATEGORIA
	r.HandleFunc(config.CATEGORIA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCategoria))).Methods(http.MethodGet)

	//	SALVA CATEGORIA
	r.HandleFunc(config.CATEGORIA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateCategoria))).Methods(http.MethodPost)

	//	EDITA CATEGORIA (cod_categoria)
	r.HandleFunc(config.CATEGORIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCategoria))).Methods(http.MethodPut)

	//	LISTA CATEGORIA (cod_categoria)
	r.HandleFunc(config.CATEGORIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCategoriaByID))).Methods(http.MethodGet)

	//	APAGA CATEGORIA (cod_categoria)
	r.HandleFunc(config.CATEGORIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteCategoria))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM CLASSE EMPENHO OK!
	=========================	*/

	//	LISTA CLASSE EMPENHO
	r.HandleFunc(config.CLASSE_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllClasseEmpenho))).Methods(http.MethodGet)

	//	SALVA CLASSE EMPENHO
	r.HandleFunc(config.CLASSE_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateClasseEmpenho))).Methods(http.MethodPost)

	//	EDITA CLASSE EMPENHO (cod_classe_empenho)
	r.HandleFunc(config.CLASSE_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateClasseEmpenho))).Methods(http.MethodPut)

	//	LISTA CLASSE EMPENHO (cod_classe_empenho)
	r.HandleFunc(config.CLASSE_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetClasseEmpenhoByID))).Methods(http.MethodGet)

	//	APAGA CLASSE EMPENHO (cod_classe_empenho)
	r.HandleFunc(config.CLASSE_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteClasseEmpenho))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ETAPA 22000
	=========================	*/

	/*	=========================
			ROTAS EM ITENS 23000
	=========================	*/

	/*	=========================
			ROTAS EM MUNICIPIOS OK!
	=========================	*/

	//	LISTA MUNICIPIOS
	r.HandleFunc(config.MUNICIPIO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllMunicipios))).Methods(http.MethodGet)

	//	SALVA MUNICIPIOS
	r.HandleFunc(config.MUNICIPIO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateMunicipios))).Methods(http.MethodPost)

	//	EDITA MUNICIPIOS (cod_ibge)
	r.HandleFunc(config.MUNICIPIO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateMunicipios))).Methods(http.MethodPut)

	//	LISTA MUNICIPIOS (cod_ibge)
	r.HandleFunc(config.MUNICIPIO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetMunicipiosByID))).Methods(http.MethodGet)

	//	APAGA MUNICIPIOS (cod_ibge)
	r.HandleFunc(config.MUNICIPIO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteMunicipios))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM NATUREZA DE DESPESA 25000
	=========================	*/

	/*	=========================
			ROTAS EM PREFEITOS OK!
	=========================	*/

	//	LISTA PREFEITOS
	r.HandleFunc(config.PREFEITOS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPrefeito))).Methods(http.MethodGet)

	//	SALVA PREFEITOS
	r.HandleFunc(config.PREFEITOS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrefeito))).Methods(http.MethodPost)

	//	EDITA PREFEITOS (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrefeito))).Methods(http.MethodPut)

	//	LISTA PREFEITOS (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrefeitoByID))).Methods(http.MethodGet)

	//	APAGA PREFEITOS (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePrefeito))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM TIPOLOGIAS 27000
	=========================	*/

	/*	=========================
			ROTAS EM TIPO ITEM 28000
	=========================	*/

	return
}
