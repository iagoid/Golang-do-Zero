package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter usuario"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values(?, ?)")
	if err != nil {
		w.Write([]byte("Falha ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Falha ao executar statement"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Falha ao obter id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprint("Usuario inserido com sucesso ", idInserido)))
}

// BuscarUsuarios: traz todos os usuarios cadastrados do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Falha ao listar usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email)
		if err != nil {
			w.Write([]byte("Falha ao escanear os usuarios"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Falha ao converter os usuarios para JSON"))
		return
	}
}

// BuscarUsuario: traz um usuario especifico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Falha ao converter os id para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Falha ao listar usuario"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		err = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email)
		if err != nil {
			w.Write([]byte("Falha ao escanear usuario"))
			return
		}
	}

	if usuario.ID == 0 {
		w.Write([]byte("Usuario não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Falha ao converter usuario para JSON"))
		return
	}
}

// AtualizarUsuario: faz o update do usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Falha ao converter os id para inteiro"))
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler campo da requisição"))
		return
	}

	var usuario usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter usuario"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Falha ao criar statement"))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(usuario.Nome, usuario.Email, ID)
	if err != nil {
		w.Write([]byte("Falha ao executar statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario: faz a exclusão do usuario no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Falha ao converter os id para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Falha ao criar statement"))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		w.Write([]byte("Falha ao executar statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
