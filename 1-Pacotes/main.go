package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Tudo o que quero utilizar pego o nome depois da ultima barra

func main() {
	fmt.Println("Escrevendo no arquivo main")
	// As funções que eu quero chamar precisam INICIAR COM LETRA MAIUSCULA
	// Porém eu posso chamar a função com LETRA MINUSCULA que está num mesmo pacote(escrever2 no auxiliar)
	// Uma boa pratica é na função importada fazer um comentário em cima dela explicando o que ela faz
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("iagoid01@gmail.com")
	fmt.Println(err)
}

/*

go mod init modulo
go build             ---> Cria um arquivo binario que contem o projeto
./modulo  			---> Roda o arquivo compilado

go get github.com/badoux/checkmail    ---> Importar modulos externos
go mod tidy    ---> Remove as dependencias não utilizadas

*/
