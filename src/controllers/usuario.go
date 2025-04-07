package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/db"
	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/model"
	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/repository"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisiao, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario model.Usuario

	if erro = json.Unmarshal(corpoRequisiao, &usuario); erro != nil {
		log.Fatal(erro)
	}

	// Agora estou abrindo a conexão com o banco de dados
	db, erro := db.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {

		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuario.ID)))

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
	w.Write([]byte("Deletando usuario..."))
}
