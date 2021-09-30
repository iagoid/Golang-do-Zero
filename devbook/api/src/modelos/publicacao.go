package modelos

import (
	"errors"
	"time"
)

// Publicacao representa a publicacao feita por um usuario
type Publicacao struct {
	ID        uint64    `json:id,omitempty`
	Titulo    string    `json:titulo,omitempty`
	Conteudo  string    `json:conteudo,omitempty`
	AutorID   uint64    `json:autorId,omitempty`
	AutorNick string    `json:autorNick,omitempty`
	Curtidas  uint64    `json:curtidas`
	CriadaEm  time.Time `json:criadaEm,omitempty`
}

// Preparar vai chamar os metodos para validar e formatar o usuario
func (usuario *Publicacao) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	return nil
}

func (usuario *Publicacao) validar(etapa string) error {
	if usuario.Titulo == "" {
		return errors.New("O titulo é obrigatório e não pode estar em branco")
	}

	if usuario.Conteudo == "" {
		return errors.New("O conteudo é obrigatório e não pode estar em branco")
	}

	return nil
}
