package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

// O texto que aparece no print será o da função closure
// pois quando a função foi declarada ela referenciava o texto de lá
func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
