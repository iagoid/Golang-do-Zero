package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Get("/categorias", GetCategorias)

	// Subroutes
	r.Route("/categorias", func(r chi.Router) {
		r.Get("/", GetCategorias)
		r.Get("/{id}", GetCategoriasByID)
	})

	r.Mount("/produtos", ProdutosRoute())

	http.ListenAndServe(":8080", r)
}

func ProdutosRoute() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	r.Get("/", func(rw http.ResponseWriter, rs *http.Request) {
		rw.Write([]byte("Produtos"))
	})

	return r
}

func GetCategorias(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Categorias"))
}

func GetCategoriasByID(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	rw.Write([]byte(id))
}
