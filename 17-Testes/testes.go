package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

/*
Executar testes sem precisar entrar no pacote
go test ./...

Rodar testes com especificações
go test -v

Quantas opções da função estão sendo cobertas pelo teste
go test --cover

Gerar relatório txt quanto da função foi coberto pelo teste
go test --coverprofile cobertura.txt
Ler relatório
go tool cover --func cobertura.txt
Visual em HTML de todas as linhas que não estão cobertas
go tool cover --html cobertura.txt
*/

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Brasil")
	fmt.Println(tipoEndereco)
}
