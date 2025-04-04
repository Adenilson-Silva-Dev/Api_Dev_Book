package routes

import (
	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/routes/rotas"
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {

	r := mux.NewRouter()

	return rotas.ConfigurarRotas(r)
}
