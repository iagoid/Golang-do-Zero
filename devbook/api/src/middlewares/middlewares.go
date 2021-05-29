package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger escreve informações sobre a rota
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s%s", r.Method, r.Host, r.RequestURI)
		next(w, r) // Executa a função que vem por parametro
	}
}

// Autenticar verifica se o usuário que faz a requisição esta autentcado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")
		next(w, r) // Executa a função que vem por parametro
	}
}
