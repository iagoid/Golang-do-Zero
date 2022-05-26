package rotas

import (
	"net/http"

	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota: representa todas as rotas da Aplicação Web
type Rota struct {
	URI                  string
	Metodo               string
	Funcao               func(http.ResponseWriter, *http.Request)
	RequerAutentificacao bool
}

// Configurar: coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasPublicacoes...)
	rotas = append(rotas, rotaLogout)

	for _, rota := range rotas {

		if rota.RequerAutentificacao {
			router.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)

		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}

	// Configur os assets para serem utilizados no template, sem a necessidade de voltar páginas
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
