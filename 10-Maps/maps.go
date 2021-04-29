package main

import "fmt"

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

}
