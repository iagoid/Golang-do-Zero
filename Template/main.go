package main

// https://coveooss.github.io/gotemplate/

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	// "text/template"
	"html/template"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoReport struct {
	Count int
	Todos []Todo
}

func main() {

	todos := RequestSite()
	reportData := TodoReport{len(todos), todos}

	TemplateSimples(&reportData)
	TemplateSimples2(&reportData)
	GenerateHTML(&reportData)
	TemplateHTML(&reportData)
	MapeamentoDeFuncao(&reportData)

}

func RequestSite() []Todo {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func TemplateSimples(reportData *TodoReport) {
	// Escrevi um template que que utilizar
	// Utilizndo a crase o texto é formtado do jeito que escrevi
	var tmpl = ` {{ .Count }} relatórios. 
		{{ range .Todos }}--------------------------------------
		ID: {{ .Id }}
		UserId: {{ .UserId }}
		Title: {{ .Title }}
		{{ if .Completed  }}Relatório Gerado{{ else }}Falha ao gerar Relatório {{ end }}
		{{ end }}
		`

	// Crio o novo template
	report, err := template.New("report").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Passo os dados para o template
	err = report.Execute(os.Stdout, reportData)
	if err != nil {
		log.Fatal(err)
	}
}

func TemplateSimples2(reportData *TodoReport) {
	// - Remove os espaços antes e depois
	/* Condicionais (operador valor1 valor 2)
	eq 		==
	ne 		!=
	lt		<
	le		<=
	gt		>
	ge		>=
	*/

	var tmpl = `
		{{- range $index, $todo := .Todos }}
		{{- if lt $index 5 }} 
			{{ $index }}: {{ $todo.Title -}}
		{{- end -}}
		{{ end }}
		`

	// Crio o novo template
	report, err := template.New("report").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Passo os dados para o template
	err = report.Execute(os.Stdout, reportData)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateHTML(reportData *TodoReport) {

	var tmpl = `
		<html>
		<head></head>
		<body>
			{{- range $index, $todo := .Todos }}
				<h1>{{ $index }}: {{ $todo.Title -}}</h1>
			{{ end }}
		</body>
		</html>
		`

	// Crio o novo template
	report, err := template.New("report").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Passo os dados para o template
	err = report.Execute(f, reportData)
	if err != nil {
		log.Fatal(err)
	}
}

func TemplateHTML(reportData *TodoReport) {
	report, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = report.Execute(f, reportData)
	if err != nil {
		log.Fatal(err)
	}
}

func MapeamentoDeFuncao(reportData *TodoReport) {

	//Atribuo a minha função quee pode ser invocada com "upper"
	// tambéem posso atribuir funções declaradas
	report, err := template.New("index2.tmpl").Funcs(template.FuncMap{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}).ParseFiles("index2.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = report.Execute(f, reportData)
	if err != nil {
		log.Fatal(err)
	}
}
