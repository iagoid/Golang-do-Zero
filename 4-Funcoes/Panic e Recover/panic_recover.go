package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada com sucesso")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// O programa não sabe o que fazer com esse valor
	// ele mata a execução do programa
	// Porém ao chamar o recover recupera a execução
	// ATENÇÃO: O panic NÃO é um ERRO
	panic("Média exatamente 6")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("-------- Programa terminou  --------")
}
