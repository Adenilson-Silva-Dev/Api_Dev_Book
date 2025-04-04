package rotas

import "net/http"

type Rotas struct {
	URI                string
	Method             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}
