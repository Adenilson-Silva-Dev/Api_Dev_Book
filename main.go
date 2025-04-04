package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adenilson-Silva-Dev/Api_Dev_Book/src/routes"
)

func main() {

	r := routes.Gerar()
	fmt.Println("Iniciando o servidor na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
