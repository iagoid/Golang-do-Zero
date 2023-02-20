package main

import (
	"embed"
	_ "embed"
	"embedded_files/templates"
	"fmt"
	"log"
	"os"
	"text/template"
)

//go:embed file.txt
var f []byte // Carrega diretaamente

//go:embed file.txt
var g embed.FS // Não carrega inicialmente, precisa ser inicializado (e testar se o arquivo pode se lido)

//go:embed templates
var t embed.FS

func main() {
	// Já carregado
	fmt.Println(string(f))

	// Não carregado (preciso testar)
	data, err := g.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Pasta
	data, err = t.ReadFile("templates/template1.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	tmpl, err := template.ParseFS(t, "templates/template1.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(os.Stdout, "R$1100,00")

	// Modo mais utilizado
	tmpl, err = template.ParseFS(templates.FS, "templates/template1.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(os.Stdout, "R$1100,00")

}
