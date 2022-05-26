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

	{
		URI:                  "/buscar-usuarios",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeUsuarios,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPerfilDoUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.PararDeSeguirUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/usuarios/{usuarioId}/seguir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.SeguirUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/perfil",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregaPerfilDoUsuarioLogado,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/editar-usuario",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeEdicaoDeUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/editar-usuario",
		Metodo:               http.MethodPut,
		Funcao:               controllers.EditarUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/atualizar-senha",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeEdicaoDeSenha,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/atualizar-senha",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarSenha,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/deletar-usuario",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarUsuario,
		RequerAutentificacao: true,
	},
}
