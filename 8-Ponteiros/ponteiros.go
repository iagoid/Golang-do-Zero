package main

import "fmt"

func main() {
	///////////////////////  Explicações ///////////////////////
	// Ponteiros são variaveis que armazenam endereços de memória
	// O & referencia o lugar na memória
	// O * desreferencial o endereço de memória
	var numero int
	var ponteiroNumero *int

	numero = 1
	ponteiroNumero = &numero
	fmt.Println("numero", numero)
	fmt.Println("ponteiroNumero", ponteiroNumero) // O valor armazenado é o endereço do numero
	fmt.Println("&numero", &numero)
	fmt.Println("&ponteiroNumero", &ponteiroNumero) // Mas o endereço é diferente
	fmt.Println("*ponteiroNumero", *ponteiroNumero) // Mostra o valor armazenado na variavel que o ponteiro aponta
	fmt.Println()
	/////////////////////// Algumas coisa ///////////////////////
	var variavel1 int = 10
	var variavel2 int = variavel1
	variavel2++
	fmt.Println(variavel1, variavel2) //10 11 pois a variavel2 copia o valor da variavel1

	var variavel3 int = 100
	var ponteiro *int = &variavel3

	*ponteiro++
	fmt.Println(variavel3, *ponteiro) //101 101 pois o ponteiro aponta para o valor do endereço da variavel3
	fmt.Println()
	////////////////////////////////// Outro exemplo //////////////////////////////////
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	fmt.Println("*pointer =", *pointer)

	// Altero o valor da variavel que o pointer aponta
	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer)

	fmt.Println("creature =", creature)

}
