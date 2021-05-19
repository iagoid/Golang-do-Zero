package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Implicitamente importado
)

/*
No mysql
create database devbook;
use devbook;
create table usuarios(
	id int auto_increment primary key,
	nome varchar(50) not null,
	email varchar(50) not null
)

*/

//go get github.com/go-sql-driver/mysql
func main() {
	stringConexao := "root:root@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão aberta")

	linhas, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
