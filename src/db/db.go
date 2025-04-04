package db

import (
	"database/sql"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/config"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoComBanco)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
