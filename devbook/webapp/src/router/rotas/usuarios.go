package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                  "/criar-usuario",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginacaoDeCadastroDeUsuario,
		RequerAutentificacao: false,
	},
	{
		URI:                  "/usuarios",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarUsuario,
		RequerAutentificacao: false,
	},
}
