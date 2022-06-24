package main

import (
	"fmt"
)

func main() {
	// ARRAYS
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(array2)

	// Apelido para o array
	type letras [5]string
	arrayLetrasInverso := letras{
		4: "A",
		3: "B",
		2: "C",
		1: "D",
		0: "E"}
	fmt.Println(arrayLetrasInverso)

	arrayLetrasSemIndice := [...]string{
		3: "C",
		"D",
		0: "A"}
	fmt.Println(arrayLetrasSemIndice)

	// Tamanho baseado na quantidade de dados que eu passar
	// NÃO POSSO ADICIONAR MAIS NUMEROS DEPOIS DA DECLARAÇÃO
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

}
