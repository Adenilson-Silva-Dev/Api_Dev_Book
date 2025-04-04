package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	URI                 string
	Method              string
	Funcao              func(w http.ResponseWriter, r *http.Request)
	RequerAuthenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {

	//Criando variaveis rotas e passando as rotas de usuarios que estao no arquivo usuario.go
	rotas := rotasUsuario

	//Percorrendo as rotas e configurando cada uma delas
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Method)

	}

	return r

}
