package main

import (
	"fmt"
)

func main() {
	array := [5]string{"A", "B", "C"}

	// SLICE
	// A diferença é não existir limite de tamanho
	// Assim posso adicionar mais dados
	// Ele é um ponteiro de um array
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	slice = append(slice, 7)
	fmt.Println(slice)

	slice2 := array[1:3]
	fmt.Println(slice2)

	array[2] = "AltereiPosição"
	fmt.Println(slice2)

	// Array interno
	fmt.Println("------------------")
	// cria um array com capacidade 11 posições e retorna um slice com 10 posicoes
	slice3 := make([]float32, 10, 11)

	fmt.Println("Tamanho", (len(slice3)), "Capacidade", (cap(slice3))) // capacidade
	fmt.Println(slice3)
	fmt.Println()

	// Ao ultrapassar a quantidade ele dobra a capacidade
	slice3 = append(slice3, 111)
	slice3 = append(slice3, 222)
	fmt.Println("Tamanho", (len(slice3)), "Capacidade", cap(slice3))
	fmt.Println(slice3)
	fmt.Println()

	// A capacidade não precisa ser especificada
	slice4 := make([]float32, 5)
	slice4 = append(slice4, 999)
	fmt.Println("Tamanho", (len(slice3)), "Capacidade", cap(slice3))
	fmt.Println(slice4)
	fmt.Println()

	// Há a possibilidade de diminuir a capacidade de um slice
	var s = []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(s[0:cap(s)])

	s = s[0:3:3] //Mudo a capacidade, porém o tamanho precisa ser igual ou menos que a capacidade
	fmt.Println(s)
	fmt.Println("Tamanho", len(s), "Capacidade", cap(s))

	// Copiar de um slice para outro (substitui se já existirem)
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	copy(s1, s2)
	fmt.Println(s1)
}
