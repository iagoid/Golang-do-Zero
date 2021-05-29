package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
