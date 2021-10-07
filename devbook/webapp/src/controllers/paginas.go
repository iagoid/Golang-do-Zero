package controllers

import (
	"net/http"
	"webapp/views/utils"
)

// CarregarTelaDeLogin: renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}

// CarregarPaginacaoDeCadastroDeUsuario: renderiza a pagina do cadastro dos usuários
func CarregarPaginacaoDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "cadastro.html", nil)
}
