package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacaos
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes: Cria um repositorio do Publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {

	statement, err := repositorio.db.Prepare(
		"insert into publicacaos (titulo, conteudo, autor_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os publicacaos que atendem um filtro de id
func (repositorio Publicacoes) BuscarPorAutor(autor_id uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		"select id, titulo, conteudo, autor_id, curtidas from publicacaos where autor_id = ?",
		autor_id)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacaos []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
		); err != nil {
			return nil, err
		}
		publicacaos = append(publicacaos, publicacao)
	}

	return publicacaos, nil
}

// BuscarPorID traz todos os publicacao que atendem um filtro de id
func (repositorio Publicacoes) BuscarPorID(id uint64) (modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		"select id, titulo, conteudo, autor_id, curtidas, criada_em from publicacaos where id = ?",
		id)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); err != nil {
			return modelos.Publicacao{}, err
		}
	}

	return publicacao, nil
}

// Atualizar atualiza um publicacao no banco de dados
func (repositorio Publicacoes) Atualizar(publicacao modelos.Publicacao) error {
	statement, err := repositorio.db.Prepare(
		"update publicacaos set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.ID)
	if err != nil {
		return err
	}

	return nil
}

// Deletar exclui um publicacao no banco de dados
func (repositorio Publicacoes) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from publicacaos where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}
