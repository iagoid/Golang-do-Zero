package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"João",
		"joao@gmail.com",
	}

	templates.ExecuteTemplate(w, "index.html", u)
}

func init() {

	templates = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
