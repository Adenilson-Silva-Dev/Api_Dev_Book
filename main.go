package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/config"
	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/db"
	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/routes"
)

func main() {

	// abrindo conexão com o banco de dados
	db.Conectar()
	// Carregar as variáveis de ambiente
	config.Carregar()
	r := routes.Gerar()
	fmt.Printf("Servidor rodando na porte %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
