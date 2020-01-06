package service

import (
	"CidadesDigitais/config"
	"CidadesDigitais/control"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateHandler() (r *mux.Router) {

	//cria um roteador
	r = mux.NewRouter()

	r.HandleFunc(config.USER_PATH, control.ListaUsuarios).Methods(http.MethodGet)
	r.HandleFunc(config.USER_PATH_LOGIN, control.Login).Methods(http.MethodPost)
	r.HandleFunc(config.USER_PATH_CREATEUSER, control.CreateUser).Methods(http.MethodPost)
	r.HandleFunc(config.USER_PATH_DELETEUSER, control.DeleteUser).Methods(http.MethodDelete)

	/*
		//**********Rotas em Usuario
					//Rotas em USER_PATH
						//LISTA USUARIO
						r.HandleFunc().Methods(http.MethodGet)
						//SALVA USUARIO
							r.HandleFunc().Methods(http.MethodPut)
						//EDITA O USUARIO
							r.HandleFunc().Methods(http.MethodPost)
					//Rotas em USER_ID_PATH
						//LISTA USUÁRIO POR ID
							r.HandleFunc().Methods(http.MethodGet)
						//APAGA O USUARIO
							r.HandleFunc().Methods(http.MethodDelete)
	*/

	/*
		//**********Rotas em Assunto
		//----------Rotas de Assunto
		//lista assunto
		r.HandleFunc().Methods(http.MethodGet)

		//salva assunto
		r.HandleFunc().Methods(http.MethodPost)

		//lista assunto por ID
		r.HandleFunc().Methods(http.MethodGet)

		//edita assunto ID
		r.HandleFunc().Methods(http.MethodPut)

		//apaga assunto ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Categoria
		//Rotas de Categoria
		//lista categotia
		r.HandleFunc().Methods(http.MethodGet)

		//salva categoria
		r.HandleFunc().Methods(http.MethodPost)

		//edita categoria
		r.HandleFunc().Methods(http.MethodPut)

		//lista categoria por ID
		r.HandleFunc().Methods(http.MethodGet)

		//deleta categoria ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Cidades Digitais
			//---------Rotas de Estado---------//
			//---------Rotas de Municipios da Cidade Digital---------//
			//---------Rotas de Cidade Digital---------//
			//---------Rotas de Itens da Cidade Digital---------//
			//---------Rotas de Processo da Cidade Digital---------//
			//---------Rotas de Acompanhamento da Cidade Digital---------//
			//---------Rotas de Ponto da Cidade Digital---------//
			//---------Rotas de Pagamento da Cidade Digital---------//

		//**********Rotas em ClasseEmpenho
		//Rotas de ClaseEmpenho
		//lista classe empenho
		r.HandleFunc().Methods(http.MethodGet)

		//salva classe empenho
		r.HandleFunc().Methods(http.MethodPost)

		//edita classe empenho
		r.HandleFunc().Methods(http.MethodPut)

		//lista classe empenho por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga classe empenho ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Contato
			//----------Rotas De Contato
				//ROTAS EM CONTATO_PATH
					//lista contato
						r.HandleFunc().Methods(http.MethodGet)

					//salva contato
						r.HandleFunc().Methods(http.MethodPost)

					//edita contato
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM CONTATO_ID_PATH
					//apaga contato ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas De Telefone
				//ROTAS EM TEL_PATH
					//lista telefone
						r.HandleFunc().Methods(http.MethodGet)

					//salva telefone
						r.HandleFunc().Methods(http.MethodPost)
				ROTAS EM TEL_ID
					//apaga telefone ID
						r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Empenho
			//---------Rotas de Empenho
				//ROTAS EM EMPENHO_PATH
					//lista empenho
						r.HandleFunc().Methods(http.MethodGet)

					//salva empenho
						r.HandleFunc().Methods(http.MethodPost)

					//edita empenho
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM EMPENHO_ID_PATH
					//lista empenho por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga emepnho ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Empenho Itens----------//
				//ROTAS EM EMPENHO_ITENS
					//edita empenho itens
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM EMPENHO_ITENS_ID
					//lista empenho itens ID
						r.HandleFunc().Methods(http.MethodGet)

		//*****************Rotas em Entidade******************
			//----------Rotas de Entidade
				//ROTAS EM ENTIDADE_PATH

					//lista entidade
						r.HandleFunc().Methods(http.MethodGet)

					//salva entidade
						r.HandleFunc().Methods(http.MethodPost)

					//edita entidade
						r.HandleFunc().Methods(http.MethodPut)

				//ROTAS EM ENTIDADE_ID_PATH

					//lista entidade por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga entidade ID
						r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Etapa
		//----------Rotas de Etapa
		//lista etapa
		r.HandleFunc().Methods(http.MethodGet)

		//salva etapa
		r.HandleFunc().Methods(http.MethodPost)

		//edita etapa
		r.HandleFunc().Methods(http.MethodPut)

		//lista etapa por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga etapa ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Fatura
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//

		//**********Rotas em Itens
		//----------Rotas de Itens
		//lista itens
		r.HandleFunc().Methods(http.MethodGet)

		//salva itens
		r.HandleFunc().Methods(http.MethodPost)

		//edita itens
		r.HandleFunc().Methods(http.MethodPut)

		//lista itens por ID (cod_item, cod_tipo_item)
		r.HandleFunc().Methods(http.MethodGet)

		//apaga itens (cod_item, cod_tipo_item)
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Lotes
			//----------Rotas de Lote
				//ROTAS EM LOTES_PATH
					//lista lote
						r.HandleFunc().Methods(http.MethodGet)

					//salva lote
						r.HandleFunc().Methods(http.MethodPost)

					//edita lote
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM LOTES_ID_PATH
					//lista lote por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga lote ID
						r.HandleFunc().Methods(http.MethodDelete)

			//---------Rotas de Reajuste
				//ROTAS EM REAJUSTE_PATH
					//lista reajuste
						r.HandleFunc().Methods(http.MethodGet)

					//salva reajuste
						r.HandleFunc().Methods(http.MethodPost)
				//ROTAS EM REAJUSTE_DEL
					//apaga reajuste (lote_cod_lote, ano_ref)
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Lote Itens
					//edita lote itens
						r.HandleFunc().Methods(http.MethodPut)

					//lista lote itens ID
						r.HandleFunc().Methods(http.MethodGet)

		//**********Rotas em Modulos
		//----------Rotas de Modulos
		//lista modulos
		r.HandleFunc().Methods(http.MethodGet)

		//lista Modulos Usuario (cod_usuario)
		r.HandleFunc().Methods(http.MethodGet)

		//**********Rotas em Municipios
		//----------Rotas de Municipio
		//lista minicipio
		r.HandleFunc().Methods(http.MethodGet)

		//salva municipio
		r.HandleFunc().Methods(http.MethodPost)

		//edita municipio
		r.HandleFunc().Methods(http.MethodPut)

		//lista municipio por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga municipio ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em NaturezaDespesa
		//----------Rotas de NaturezaDespesa
		//lista natureza despesa
		r.HandleFunc().Methods(http.MethodGet)

		//salva natureza despesa
		r.HandleFunc().Methods(http.MethodPost)

		//edita natureza despesa
		r.HandleFunc().Methods(http.MethodPut)

		//lista natureza despesa por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga natureza despesa ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Pagamento (otb)
			//----------Rotas de Pagamento
				//ROTAS EM PAG_PATH
					//lista pagamento
						r.HandleFunc().Methods(http.MethodGet)

					//salva pagamento
						r.HandleFunc().Methods(http.MethodPost)

					//edita pagamento
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM PAG_ID_PATH
					//lista pagamento por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga pagamento ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Municipio Fatura
				//ROTAS EM PAG_FAT_MUNICIPIO
					//lista municipio fatura (cd_municipio_cod_ibge)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_FAT_LIST
					//lista pagamento fatura (cod_otb)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_FAT_SAVE
					//salva pagamento fatura
						r.HandleFunc().Methods(http.MethodPost)

			//----------Rotas de Pagamento Itens
				//ROTAS EM PAG_LIST_ITENS
					//lista pagamento itens	(cod_otb)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_EDIT_ITENS
					//edita pagamento itens
						r.HandleFunc().Methods(http.MethodPut)

		//**********Rotas em Prefeitos
		//----------Rotas de Prefeitos
		//lista prefeitos
		r.HandleFunc().Methods(http.MethodGet)

		//salva prefeitos
		r.HandleFunc().Methods(http.MethodPost)

		//edita prefeitos
		r.HandleFunc().Methods(http.MethodPut)

		//lista prefeitos por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga prefeitos ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em PrevisaoEmpenho
		//----------Rotas de PrevisaoEmpenho
		//lista previsao empenho
		r.HandleFunc().Methods(http.MethodGet)

		//salva previsao empenho
		r.HandleFunc().Methods(http.MethodPost)

		//edita previsao empenho
		r.HandleFunc().Methods(http.MethodPut)

		//lista previsao emprenho por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga previsao empenho ID
		r.HandleFunc().Methods(http.MethodDelete)

		//----------Rotas de PrevisaoEmpenhoItens
		//edita previsao empenho itens
		r.HandleFunc().Methods(http.MethodPut)

		//lista previsao empenho itens (cod_lote, cod_previsao_empenho)
		r.HandleFunc().Methods(http.MethodGet)

		//**********Rotas em TipoItem

		//**********Rotas em Tipologia



	*/
	return
}

//Methods: OPTIONS, GET, POST, PUT, DELETE
