package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 5 {
		i++
		fmt.Println("Incremento i: ", i)
		time.Sleep(time.Second)
	}

	for j := 1; j <= 5; j++ {
		fmt.Println("Incremento j: ", j)
		time.Sleep(time.Second)
	}

	// Arrays
	nomes := [3]string{"Iago", "Lucas", "Pedro"}

	fmt.Println("\nIndice")
	for indice := range nomes {
		fmt.Println(indice)
	}

	fmt.Println("\nIndices e Nomes")
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("\nEm palavras")
	for indice, letra := range "PALAVRA" {
		// Retorna pela tabela ASCII
		fmt.Println(indice, "ASCII -->", letra, "=", string(letra))
	}

	// MAPs
	fmt.Println("\nMAP")
	usuario := map[string]string{
		"nome":      "Cleiton",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// LOOP INFINITO
	for {
		fmt.Println("LOOP INFINITO")
		time.Sleep(time.Second)
	}

}
