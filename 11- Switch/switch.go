package main

import "fmt"

// NÃO É NECESSÁRIO UTILIZAR BREAK
func dias(numero int) string {
	switch numero {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda"

	case 3:
		return "Terça"

	case 4:
		return "Quarta"

	case 5:
		return "Quinta"

	case 6:
		return "Sexta"

	case 7:
		return "Sábado"

	default:
		return "Dia não identificado"
	}
}

// Quando não quero avaliar a mesma variavel
func diaDaSemana(numero int) string {
	switch {
	case numero == 2:
		return "Segunda"

	case numero == 3:
		return "Terça"

	case numero == 4:
		return "Quarta"

	case numero == 5:
		return "Quinta"

	case numero == 6:
		return "Sexta"

	default:
		return "Dia não identificado"
	}
}

func fimDeSemana(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough // Pega a proxima condição

	case numero == 7:
		dia = "Sábado"
		fallthrough // Pega a proxima condição

	default:
		dia = "Dia não identificado"
	}
	return dia
}

func main() {
	// NÃO É NECESSÁRIO UTILIZAR BREAK
	dia := dias(5)
	fmt.Println(dia)

	dia2 := diaDaSemana(2)
	fmt.Println(dia2)

	dia3 := fimDeSemana(7)
	fmt.Println(dia3)
}
