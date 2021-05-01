package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	// Em casos que eu não passo argumentos, ele pega os valores
	// padrão definidos nas flags do app (host e google.com.br)

	// go run main.go NAME_COMMANDS --NAME_FLAGS SITE
	// go run main.go ip --host google.com.br
	// go run main.go servidores --host google.com.br
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
