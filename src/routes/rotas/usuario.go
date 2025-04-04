package rotas

import (
	"net/http"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/controllers"
	"github.com/gorilla/mux"
)

var rotasUsuario = []Rotas{
	{
		URI:                "/usuario",
		Method:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario",
		Method:             http.MethodGet,
		Funcao:             controllers.BuscandoUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario/usuario{id}",
		Method:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario/usuario{id}",
		Method:             http.MethodPut,
		Funcao:             controllers.AtualizandoUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuario/usuario{id}",
		Method:             http.MethodDelete,
		Funcao:             controllers.DeletandoUsuario,
		RequerAutenticacao: false,
	},
}

func ConfigurarRotas(r *mux.Router) *mux.Router {

	//Criando uma variavel rotas e passando as rotas de usuario para que eu possa iterar sobre elas
	rotas := rotasUsuario


	// Irei irerar em cada rota e passar o metodo e a funcao
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Method)
	}

	return r

}
