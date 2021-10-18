package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger: escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s%s", r.Method, r.Host, r.RequestURI)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, err := cookies.Ler(r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		proximaFuncao(w, r)
	}
}
