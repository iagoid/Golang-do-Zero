package main

import (
	"fmt"
	"sort"
)

func main() {
	// O map é muito dependente do tipo que ele foi declarado
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// Map aninhado com outro map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Iago",
			"ultimo":   "Dalmolin",
		},
		"curso": {
			"nome":   "CC",
			"campus": "Ibiruba",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])

	// Deletando chave
	delete(usuario2, "curso")

	// Adicionando chave
	usuario2["signo"] = map[string]string{
		"nome": "escorpião",
	}

	fmt.Println(usuario2)

	//////////////////////////////////////////////////
	// Ordenar map
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	// pega os valores do map
	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Ordena os valores
	sort.Ints(values)

	for _, v := range values {
		fmt.Print(v, " - ")
	}

	// Ordena as chaves

	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}

}
