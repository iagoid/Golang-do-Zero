package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Retornos multiplos
func calculosMatematicos(n1 int8, n2 int8) (int8, int8, int8, int8) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

func main() {
	soma := somar(10, 15)
	fmt.Println(soma)

	// Função na variável
	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("TEXTO NA FUNÇÃO F")

	// Caso eu não queira utilizar o valor eu coloco _
	_, subtracao, multiplicacao, divisao := calculosMatematicos(10, 15)
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
}
