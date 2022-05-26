package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:                  "/publicacoes",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarPublicacao,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/publicacoes/{publicacaoId}/curtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CurtirPublicacao,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/publicacoes/{publicacaoId}/descurtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.DescurtirPublicacao,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/publicacoes/{publicacaoId}/atualizar",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeEdicaoDePublicacao,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/publicacoes/{publicacaoId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarPublicacao,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/publicacoes/{publicacaoId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarPublicacao,
		RequerAutentificacao: true,
	},
}
