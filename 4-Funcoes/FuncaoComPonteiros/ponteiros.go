package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

// Altera o numero da variavel
// Não necessita de retorno pois eu altero direto na memoria
func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println("Invertido", numeroInvertido)
	fmt.Println("Numero", numero)

	novoNumero := 20
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println("Invertido", novoNumero)
	fmt.Println("Numero", novoNumero)
}
