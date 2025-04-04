package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoComBanco = ""
	Porta                 = 0
)

// funcão  vai ser responsável por carregar as variáveis de ambiente
func Carregar() {

	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// convertendo a porta para uma string usando o strconv.Atoi

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	StringConexaoComBanco = fmt.Sprintf("%s:%s@tcp(localhost:3305)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
