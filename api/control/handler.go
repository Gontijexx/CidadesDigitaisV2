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
			ROTAS EM CD
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
			ROTAS EM CD_ITENS
	=========================	*/

	//	LISTA CD_ITENS
	r.HandleFunc(config.CD_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCDItens))).Methods(http.MethodGet)

	//	EDITA CD_ITENS
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCDItens))).Methods(http.MethodPut)

	//	LISTA CD_ITENS POR ID
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDItensByID))).Methods(http.MethodGet)

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

	//	APAGA REAJUSTE (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePrevisaoEmpenho))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ITENS PREVISAO EMPENHO OK!
	=========================	*/

	//	EDITA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensPrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS PREVISAO EMPENDO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensPrevisaoEmpenhoByID))).Methods(http.MethodGet)

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
			ROTAS EM PREFEITOS OK!
	=========================	*/

	//	LISTA EMPENHO
	r.HandleFunc(config.PREFEITOS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPrefeito))).Methods(http.MethodGet)

	//	SALVA EMPENHO
	r.HandleFunc(config.PREFEITOS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrefeito))).Methods(http.MethodPost)

	//	EDITA EMPENHO (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrefeito))).Methods(http.MethodPut)

	//	LISTA EMPENHO (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrefeitoByID))).Methods(http.MethodGet)

	//	APAGA EMPENHO (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePrefeito))).Methods(http.MethodDelete)

	return
}
