package rotas

import (
	"net/http"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/controllers"
)

var rotasUsuario = []Rotas{
	{
		URI:                 "/usuario",
		Method:              http.MethodPost,
		Funcao:              controllers.CriarUsuario,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuario",
		Method:              http.MethodGet,
		Funcao:              controllers.BuscandoUsuarios,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuario/usuario{id}",
		Method:              http.MethodGet,
		Funcao:              controllers.BuscandoUsuarios,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuario/usuario{id}",
		Method:              http.MethodPut,
		Funcao:              controllers.AtualizandoUsuario,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuario/usuario{id}",
		Method:              http.MethodDelete,
		Funcao:              controllers.DeletandoUsuario,
		RequerAuthenticacao: false,
	},
}
