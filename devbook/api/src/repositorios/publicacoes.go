package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes: Cria um repositorio do Publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar: insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {

	statement, err := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
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

// BuscarPorID: traz todos os publicacao que atendem um filtro de id
func (repositorio Publicacoes) BuscarPorID(id uint64) (modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		`select p.id, titulo, conteudo, autor_id, curtidas, criadaEm, u.nick from publicacoes p
		inner join usuarios u ON u.id = p.autor_id
		where p.id = ?`,
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
			&publicacao.AutorNick,
		); err != nil {
			return modelos.Publicacao{}, err
		}
	}

	return publicacao, nil
}

// Atualizar: atualiza um publicacao no banco de dados
func (repositorio Publicacoes) Atualizar(publicacao modelos.Publicacao) error {
	statement, err := repositorio.db.Prepare(
		"update publicacoes set titulo = ?, conteudo = ? where id = ?")
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

// Deletar: exclui um publicacao no banco de dados
func (repositorio Publicacoes) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from publicacoes where id = ?")
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

// BuscarPorAutor: traz todos os publicacaos dos usuarios que atendem um filtro de id
func (repositorio Publicacoes) BuscarPorAutor(autor_id uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		`select p.id, titulo, conteudo, autor_id, curtidas, criadaEm, u.nick from publicacoes p
		inner join usuarios u ON u.id = p.autor_id
		where u.id = ?`,
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
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacaos = append(publicacaos, publicacao)
	}

	return publicacaos, nil
}

// BuscarPublicacoes: traz todos os publicacaos que atendem um filtro de id
func (repositorio Publicacoes) BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		`select distinct p.*, u.nick 
		from publicacoes p
		inner join usuarios u ON u.id = p.autor_id
		left join seguidores s ON p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id = ?
		order by 1`,
		usuarioID, usuarioID)
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
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacaos = append(publicacaos, publicacao)
	}

	return publicacaos, nil
}

// Curtir: aumenta o numero de curtidas na publicacao
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

// Descurtir: subtrai o numero de curtidas na publicacao
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare(`update publicacoes set curtidas = 
			CASE WHEN curtidas > 0 THEN  curtidas - 1
			ELSE curtidas END
			where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
