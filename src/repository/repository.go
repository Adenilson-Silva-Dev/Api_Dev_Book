package repository

import (
	"database/sql"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/model"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um novo repositorio de usuario
func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {

	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario model.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha)values(?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	// passando dessa linha, signifia que o usuario já foi ccadastrado do baco
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro !=  nil{
		return 0, erro
	}

	ultimoIDInserido, erro:= resultado.LastInsertId()

	if erro != nil{
		return 0,erro
	}

	// Convertendo o ultimoIDInserido para uint64
	return uint64(ultimoIDInserido), nil

}
