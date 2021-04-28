package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

// "Herda" da pessoa e joga para dentro de estudante
type estudante struct {
	pessoa // Permite pegar um dado com mais facilidade
	curso  string
}

func main() {
	pessoa := pessoa{"Iago", 19}

	estudante := estudante{pessoa, "CC"}
	fmt.Println(estudante)
	fmt.Println(estudante.nome)
	fmt.Println(estudante.idade)
	fmt.Println(estudante.curso)
}
