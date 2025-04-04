package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario..."))
}
func BuscandoUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuario..."))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario..."))
}
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario..."))
}
func DeletandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario..."))
}
