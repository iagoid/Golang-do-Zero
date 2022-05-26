package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa uma pessoa acessando a aplicação
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicaacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuario
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	var usuario Usuario

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar usuario")
			}
			usuario = usuarioCarregado

		case seguidoresCarregado := <-canalSeguidores:
			if seguidoresCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar seguidores")
			}
			usuario.Seguidores = seguidoresCarregado

		case seguidosCarregado := <-canalSeguindo:
			if seguidosCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar seguidos")
			}
			usuario.Seguindo = seguidosCarregado

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar as publicacoes")
			}
			usuario.Publicacoes = publicacoesCarregadas
		}
	}

	return usuario, nil
}

// BuscarDadosDoUsuario chama a API para buscar os dados base do usuario
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		canal <- Usuario{}
	}

	canal <- usuario
}

// BuscarSeguidores chama a API para buscar os seguidores do usuario
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []Usuario{}
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if err = json.NewDecoder(response.Body).Decode(&seguidores); err != nil {
		canal <- []Usuario{}
	}

	// Ocorre um erro no template quando o seguidores estão vazios
	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

// BuscarSeguindo chama a API para buscar os usuarios que o usuario do id segue
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []Usuario{}
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if err = json.NewDecoder(response.Body).Decode(&seguindo); err != nil {
		canal <- []Usuario{}
	}

	// Ocorre um erro no template quando o seguidos estão vazios
	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

// BuscarPublicacoes chama a API para buscar as publicacoes do usuario
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes/autor/%d", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []Publicacao{}
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		canal <- []Publicacao{}
	}

	// Ocorre um erro no template quando as publicações estão vazios
	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
