package main

import "fmt"

var variavel7 = "Variavel de escopo do pacote"

func main() {
	// Duas maneiras de declarar uma variavel
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Varias de uma vez
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3 + variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5 + variavel6)

	// Tudo realizado anteriormente pode ser feito com constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Invertendo valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5 + variavel6)

	// Quando quero mostrar ummesmo valor varias vezes no print
	fmt.Printf("%v %v --- %[1]v %[2]v\n", variavel1, variavel2)

	fmt.Println(variavel7)
	var variavel7 = "Variavel 7"
	fmt.Println(variavel7) //shadowing

}
