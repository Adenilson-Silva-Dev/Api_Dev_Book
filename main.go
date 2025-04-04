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
	fmt.Println(config.Porta)
	r := routes.Gerar()
	fmt.Println("Iniciando o servidor na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
