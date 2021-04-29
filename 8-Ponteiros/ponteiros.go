package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	variavel2++
	fmt.Println(variavel1, variavel2)

	//Referencia de memória
	var variavel3 int = 100
	// O & referencia o lugar na memória
	var ponteiro *int = &variavel3

	// O * desreferencial o endereço de memória
	*ponteiro++
	fmt.Println(variavel3, *ponteiro)

}
