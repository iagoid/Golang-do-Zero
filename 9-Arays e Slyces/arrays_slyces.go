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

	// Tamanho baseado na quantidade de dados que eu passar
	// NÃO POSSO ADICIONAR MAIS NUMEROS DEPOIS DA DECLARAÇÃO
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICE
	// A diferença é não existir limite de tamanho
	// Assim posso adicionar mais dados
	// Ele é um ponteiro de um array
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	slice = append(slice, 7)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[2] = "AltereiPosição"
	fmt.Println(slice2)

	// Array interno
	fmt.Println("------------------")
	// cria um array com capacidade 11 posições e retorna um slice com 10 posicoes
	slice3 := make([]float32, 10, 11)

	fmt.Println((len(slice3))) // tamanho
	fmt.Println((cap(slice3))) // capacidade
	fmt.Println(slice3)
	fmt.Println()

	// Ao ultrapassar a quantidade ele dobra a capacidade
	slice3 = append(slice3, 111)
	slice3 = append(slice3, 222)
	fmt.Println((len(slice3)))
	fmt.Println((cap(slice3)))
	fmt.Println(slice3)
	fmt.Println()

	// A capacidade não precisa ser especificada
	slice4 := make([]float32, 5)
	slice4 = append(slice4, 999)
	fmt.Println((len(slice4)))
	fmt.Println((cap(slice4)))
	fmt.Println(slice4)

}
