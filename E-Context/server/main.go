package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Verifica se a request foi cancelada antes dos 5 segundos
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Inicio da request")
	defer log.Println("Finalizou a request")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Requisição processada com sucesso")
		w.Write([]byte("Requisição processada com sucesso"))
	case <-ctx.Done():
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}
