package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("--------- Media calculada ---------")

	fmt.Println("Entrando na função de medias")

	media := (n1 + n2) / 2
	if media >= 7 {
		return true
	}
	return false
}

/*
	DEFER adia a execução da função até o ultimo momento possivel
	(no caso o fim da função main)
	Em casos de funçoes com return ele executão a função adiada antes do retorno
*/
func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(7, 8))
	fmt.Println("---------- Acabou o main ----------")
}
