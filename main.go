package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/config"
	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/routes"
)

func main() {

	// Carregar as variáveis de ambiente
	config.Carregar()
	r := routes.Gerar()
	fmt.Printf("Servidor rodando na porte %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
