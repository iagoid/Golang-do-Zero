package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variavel templates
func CarregarTemplates() {
	// ATENÇÃO: O goolang para de renderizar no primeiro erro no template
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplates executa os dados da tela
func ExecutarTemplates(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
